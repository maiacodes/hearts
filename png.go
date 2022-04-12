package main

import (
	"bytes"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
	"image"
	"image/png"
	"os"
)

func writePNG(data, name string) error {
	w, h := 512, 512
	icon, _ := oksvg.ReadIconStream(bytes.NewBufferString(data))
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds())), 1)

	out, err := os.Create(name + ".png")
	if err != nil {
		return err
	}
	defer out.Close()

	err = png.Encode(out, rgba)
	if err != nil {
		return err
	}

	return nil
}
