package main

import (
	. "github.com/eyedeekay/go-ccw"
	"log"
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
	CHROMIUM, ERROR = SecureExtendedChromium("i2pchromium", false, EXTENSIONS, EXTENSIONHASHES, ARGS...)
	if ERROR != nil {
		log.Fatal(ERROR)
	}
	defer CHROMIUM.Close()
	<-CHROMIUM.Done()
}
