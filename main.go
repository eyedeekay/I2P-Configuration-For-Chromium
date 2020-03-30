//go:generate go run -tags generate gen.go

package main

import (
	"io/ioutil"
	"log"
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

func main() {
	if embedded, err := FS.Readdir(0); err != nil {
		log.Println("Extension error, embedded extension not read.")
	} else {
		if _, err := os.Stat("i2pchrome.js"); os.IsNotExist(err) {
			os.MkdirAll("i2pchrome.js", FS.Mode())
			for _, val := range embedded {
				//log.Println(val.Name())
				ioutil.WriteFile("i2pchrome.js"+val.Name(), val.Sys().([]byte), val.Mode())
			}
		} else {
			log.Println("i2pchrome plugin already found")
		}
	}
	CHROMIUM, ERROR = SecureExtendedChromium("i2pchromium", false, EXTENSIONS, EXTENSIONHASHES, ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer CHROMIUM.Close()
	<-CHROMIUM.Done()
}
