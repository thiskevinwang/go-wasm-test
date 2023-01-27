.DEFAULT_GOAL := default

copy_wasm_exec:
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./html

build:
	pushd ./go && GOOS=js GOARCH=wasm go build -o ../html/main.wasm main.go

serve:
	npx serve ./html

default: build serve

