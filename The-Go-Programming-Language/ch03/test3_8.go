package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	getImage("complex64")
	getImage("complex128")
	getImage("bigfloat")
	getImage("bigrat")
}

func getImage(filename string) {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			var c color.Color

			y_ := float64(py)/height*(ymax-ymin) + ymin
			x_ := float64(px)/width*(xmax-xmin) + xmin
			switch filename {
			case "complex64":
				c = mandelbrotC64(z_)

			case "complex128":
				z_ := complex(x_, y_)
				c = mandelbrotC128(z_)

			case "bigfloat":
				z_ := complex(x_, y_)
				c = mandelbrotC128(z_)

			case "bigrat":
				z_ := complex(x_, y_)
				c = mandelbrotC128(z_)
			}
			// Image point (px, py) represents complex value z.
			img.Set(px, py, c)
		}
	}
	f, err := os.Create(filename + ".png")
	if err != nil {
		panic(err)
	}
	png.Encode(f, img)
	f.Close()
}

func mandelbrotC64(z complex64) color.Color {
}

func mandelbrotC128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15 //对比度
	const e = 1e-6
	// 牛顿迭代法求解z^4=1
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
