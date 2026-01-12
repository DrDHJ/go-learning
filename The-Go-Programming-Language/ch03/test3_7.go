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
					c := mandelbrotNiu(z_)
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
	f, err := os.Create("mandelbrotNiu.png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
	f.Close()
}

func mandelbrotNiu(z complex128) color.Color {
	const iterations = 200
	const contrast = 15 //对比度
	const e = 1e-6

	v := z
	for n := uint8(0); n < iterations; n++ {
		v = v - (cmplx.Pow(v, 4)-1)/(4*cmplx.Pow(v, 3))
		if cmplx.Abs(v-1) < e {
			return color.RGBA{255 - n*contrast, 0, 0, 255}
		} else if cmplx.Abs(v+1) < e {
			return color.RGBA{0, 255 - n*contrast, 0, 255}
		} else if cmplx.Abs(v-complex(0, 1)) < e {
			return color.RGBA{0, 0, 255 - n*contrast, 255}
		} else if cmplx.Abs(v-complex(0, -1)) < e {
			return color.RGBA{255 - n*contrast, 255 - n*contrast, 0, 255}
		}
	}
	return color.Black
}
