package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Call("alert", "hello")
	fmt.Println("Hello, WebAssembly!")
}
