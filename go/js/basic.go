//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"
)

func main() {
	g := js.Global()

	g.Call("alert", "hello")
	g.Call("updatedHandler", "Wasm Go JS demo")

	doc := g.Get("document")
	h2 := doc.Call("getElementsByTagName", "h2").Call("item", "0")
	h2.Set("innerHTML", "Dipankar")
}
