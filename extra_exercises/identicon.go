package main

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
)

func main() {
	var buf bytes.Buffer
	m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	
    png.Encode(&buf, m)
	ioutil.WriteFile("image.png", buf.Bytes(), 0644)

}
