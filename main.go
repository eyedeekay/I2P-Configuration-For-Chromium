//go:generate go run -tags generate gen.go

package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	. "github.com/eyedeekay/go-ccw"
)

var EXTENSIONS = []string{"i2pchrome.js"}
var EXTENSIONHASHES = []string{"359023d7c0e3eff50797c39942b27d088bd6db70740374dc3cf547fa540328f4"}
var ARGS = []string{
	"--safebrowsing-disable-download-protection",
	"--disable-client-side-phishing-detection",
	"--disable-3d-apis",
	"--disable-accelerated-2d-canvas",
	"--disable-remote-fonts",
	"--disable-sync-preferences",
	"--disable-sync",
	"--disable-speech",
	"--disable-webgl",
	"--disable-reading-from-canvas",
	"--disable-gpu",
	"--disable-32-apis",
	"--disable-auto-reload",
	"--disable-background-networking",
	"--disable-d3d11",
	"--disable-file-system",
}

func writeSubDirectory(fs http.File) {
	log.Println("writing subdirectory")
	name, err := fs.Stat()
	if err != nil {
		log.Fatal(err)
	}
	if embedded, err := fs.Readdir(0); err != nil {
		log.Println("Extension error, embedded extension not read.")
	} else {
		if _, err := os.Stat("i2pchrome.js"); os.IsNotExist(err) {
			os.MkdirAll("i2pchrome.js/"+name.Name(), FS.Mode())
			for _, val := range embedded {
				file, err := FS.Open(val.Name()) //
				if err != nil {
					log.Fatal(err.Error())
				}
				sys := bytes.NewBuffer(nil)
				if _, err := io.Copy(sys, file); err != nil {
					log.Fatal(err.Error())
				}
				ioutil.WriteFile("i2pchrome.js/"+name.Name()+"/"+val.Name(), sys.Bytes(), val.Mode())
			}
		} else {
			log.Println("i2pchrome plugin already found")
		}
	}
}

func writeExtension(val os.FileInfo, system http.FileSystem) {
	if len(val.Name()) > 3 {
		if val.IsDir() {
			os.MkdirAll("i2pchrome.js/"+val.Name(), FS.Mode())
			file, err := FS.Open(val.Name()) //
			if err != nil {
				log.Fatal(err.Error())
			}
			writeSubDirectory(file)
		} else {
			log.Println("Writing file to extension", val.Name())
			file, err := FS.Open(val.Name()) //
			if err != nil {
				log.Fatal(err.Error())
			}
			sys := bytes.NewBuffer(nil)
			if _, err := io.Copy(sys, file); err != nil {
				log.Fatal(err.Error())
			}
			if err := ioutil.WriteFile("i2pchrome.js/"+val.Name(), sys.Bytes(), val.Mode()); err != nil {
				log.Fatal(err.Error())
			}
		}
	} else {
		log.Println("+i2pchrome.js/"+val.Name()+"'", "ignored", "contents", val.Sys())
	}
}

func writeProfile(system http.FileSystem) {
	if embedded, err := FS.Readdir(0); err != nil {
		log.Println("Extension error, embedded extension not read.")
	} else {
		if _, err := os.Stat("i2pchrome.js"); os.IsNotExist(err) {
			os.MkdirAll("i2pchrome.js/icons", FS.Mode())
			os.MkdirAll("i2pchrome.js/options", FS.Mode())
			os.MkdirAll("i2pchrome.js/_locales/en", FS.Mode())
			for _, val := range embedded {
				writeExtension(val, FS)
			}
		} else {
			log.Println("i2pchrome plugin already found")
		}
	}
}

func main() {
	writeProfile(FS)
	CHROMIUM, ERROR = SecureExtendedChromium("i2pchromium-browser", false, EXTENSIONS, EXTENSIONHASHES, ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer CHROMIUM.Close()
	<-CHROMIUM.Done()
}
