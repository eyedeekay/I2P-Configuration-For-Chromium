package main

import (
    "log"
    . "github.com/eyedeekay/go-ccw"
)

var EXTENSIONS = []string{"./i2pchrome.js"}

var CHROMIUM, ERROR = ExtendedChromium("basic", true, EXTENSIONS, "")

func main() {
    if err := Run(); err != nil {
        log.Fatal(err)
    }
}