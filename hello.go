package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	// Create a channel which prevents the application from exiting
	done := make(chan struct{}, 0)

	// Create a callback function to be called from javascript
	// Simply prints Hello, world! to stdout
	cb := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		fmt.Println("Hello, world!")
		return nil
	})

	// Add click event listener to the button that will execute the callback function
	js.Global().Get("document").Call("getElementById", "myButton").Call("addEventListener", "click", cb)

	<-done
}
