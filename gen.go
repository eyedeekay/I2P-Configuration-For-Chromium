//+build generate

package main

import "github.com/zserge/lorca"

func main() {
	// You can also run "npm build" or webpack here, or compress assets, or
	// generate manifests, or do other preparations for your assets.
	lorca.Embed("main", "assets.go", "i2pchrome.js")
	lorca.Embed("i2pchrome", "lib/assets.go", "i2pchrome.js")
}
