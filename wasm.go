package main

import (
	"syscall/js"
	"time"
)

func main() {
	document := js.Global().Get("document")

	button := document.Call("getElementById", "myButton")
	button.Call("addEventListener", "click", js.NewCallback(func(args []js.Value) {
		println("Button clicked!")
	}))
	println("Hello, WebAssembly!")

	time.Sleep(time.Second * 300)
}
