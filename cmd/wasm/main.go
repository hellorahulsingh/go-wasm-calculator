package main

import (
	"strconv"
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", args[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", args[2].String()).Set("value", int1+int2)

	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	value1 := js.Global().Get("document").Call("getElementById", args[0].String()).Get("value").String()
	value2 := js.Global().Get("document").Call("getElementById", args[1].String()).Get("value").String()

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)

	js.Global().Get("document").Call("getElementById", args[2].String()).Set("value", int1-int2)

	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	registerCallbacks()
	<-c
}
