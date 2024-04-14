//go:build js && wasm
// +build js,wasm

package main

import (
	"bytes"
	"image"
	"image/color"
	"image/jpeg"
	"log/slog"
	"syscall/js"
)

func main() {
	ch := make(chan struct{})
	exports()
	<-ch
}

func exports() {
	js.Global().Set("grayscale", js.FuncOf(grayscale))
}

func grayscale(this js.Value, args []js.Value) any {
	input := make([]byte, args[0].Get("length").Int())

	js.CopyBytesToGo(input, args[0])

	img, err := jpeg.Decode(bytes.NewReader(input))
	if err != nil {
		slog.Error("Error decoding", "err", err)
	}

	bounds := img.Bounds()
	gray := image.NewGray(bounds)

	slog.Info("debugging", "input", input, "image", img, "bounds", bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			px := img.At(x, y)
			gr := color.GrayModel.Convert(px)
			gray.Set(x, y, gr)
		}
	}

	var buf bytes.Buffer
	jpeg.Encode(&buf, gray, nil)
	jsOut := js.Global().Get("Uint8Array").New(len(buf.Bytes()))
	js.CopyBytesToJS(jsOut, buf.Bytes())
	slog.Info("Uint8Array", "jsOut", jsOut, "buf", buf.Bytes())
	return jsOut
}
