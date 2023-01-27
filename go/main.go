// Package must be named `main` or else you will get a
// "Expected magic word" error in the console when trying to instantiate the wasm
// - "https://github.com/golang/go/issues/35657"
package main

import (
	"crypto"
	_ "crypto/sha512"
	"encoding/hex"
	"strconv"
	"syscall/js"

	"github.com/google/uuid"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("wasmHash", js.FuncOf(hash))
	js.Global().Set("wasmNanoid", js.FuncOf(nanoid))
	js.Global().Set("wasmUuid", js.FuncOf(gen_uuid))
	<-done
}

func hash(this js.Value, args []js.Value) interface{} {
	h := crypto.SHA512.New()
	h.Write([]byte(args[0].String()))

	return hex.EncodeToString(h.Sum(nil))
}

func nanoid(this js.Value, args []js.Value) interface{} {
	input := args[0].String()
	intVar, _ := strconv.Atoi(input)
	id, _ := gonanoid.New(intVar)
	return id
}

func gen_uuid(this js.Value, args []js.Value) interface{} {
	return uuid.New().String()
}
