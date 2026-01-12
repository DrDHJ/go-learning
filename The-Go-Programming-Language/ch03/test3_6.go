package main

import (
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
		scale                  = 2
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var r, g, b, a uint8
			for i := 0; i < scale; i++ {
				for j := 0; j < scale; j++ {
					y_ := (float64(py)+float64(i)/float64(scale))/height*(ymax-ymin) + ymin
					x_ := (float64(px)+float64(j)/float64(scale))/width*(xmax-xmin) + xmin

					z_ := complex(x_, y_)
					c := mandelbrot2(z_)
					cr, cg, cb, ca := c.RGBA()
					r += uint8(cr) / uint8(scale*scale)
					g += uint8(cg) / uint8(scale*scale)
					b += uint8(cb) / uint8(scale*scale)
					a += uint8(ca) / uint8(scale*scale)
				}
			}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, color.RGBA{r, g, b, a})
		}
	}
	f, err := os.Create("mandelbrot2.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
	f.Close()
}

func mandelbrot2(z complex128) color.Color {
	const iterations = 200
	const contrast = 15 //对比度

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			// t := float64(n) / float64(iterations)
			// r := uint8(9 * (1 - t) * t * t * t * 255)
			// g := uint8(15 * (1 - t) * (1 - t) * t * t * 255)
			// b := uint8(8.5 * (1 - t) * (1 - t) * (1 - t) * t * 255)
			// return color.RGBA{r, g, b, 255}
			return color.RGBA{255 - n*8, 222 - n*4, n * 2, 255}
		}
	}
	return color.Black
}
