goos=wasip1
goarch=wasm

build:
	GOOS=$(goos) GOARCH=$(goarch) go build -o run.wasm .

run:
	wasmtime run.wasm
