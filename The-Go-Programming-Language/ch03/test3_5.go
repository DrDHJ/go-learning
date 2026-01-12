package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrotWithColor(z))
		}
	}
	f, err := os.Create("mandelbrotWithColor.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create file: err", err)
		os.Exit(1)
	}
	err = png.Encode(f, img)
	if err != nil {
		fmt.Fprintf(os.Stderr, "encode png: err", err)
		os.Exit(1)
	}
	f.Close()
}

func mandelbrotWithColor(z complex128) color.Color {
	const iterations = 200
	const contrast = 15 //对比度

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			t := float64(n) / float64(iterations)
			r := uint8((1 - t) * 255)
			g := uint8((1 - t) * t * 255 * 4)
			b := uint8(t * 255)
			// 另一种颜色方案
			// r := uint8(9 * (1 - t) * t * t * t * 255)
			// g := uint8(15 * (1 - t) * (1 - t) * t * t * 255)
			// b := uint8(8.5 * (1 - t) * (1 - t) * (1 - t) * t * 255)
			return color.RGBA{r, g, b, 255}

			// return color.RGBA{255 - contrast*n, 0, contrast * n, 255}
		}
	}
	return color.Black
}
