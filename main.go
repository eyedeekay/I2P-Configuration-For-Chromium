package main

import (
    "log"
    . "github.com/eyedeekay/go-ccw"
)

var EXTENSIONS = []string{"./i2pchrome.js"}

var CHROMIUM, ERROR = ExtendedChromium("basic", true, EXTENSIONS, "")

func main() {
    if ERROR != nil {
        log.Fatal(ERROR)
    }
    defer CHROMIUM.Close()
    <-CHROMIUM.Done()
}