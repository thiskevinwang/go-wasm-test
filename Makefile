.DEFAULT_GOAL := default

copy_wasm_exec:
	cp $$(go env GOROOT)/misc/wasm/wasm_exec.js ./docs

build:
	pushd ./go && GOOS=js GOARCH=wasm go build -o ../docs/main.wasm main.go

serve:
	npx serve ./docs

default: build serve

