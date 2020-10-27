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
		width, height = 1024, 1024
		pngfile = "mandelbrot.png"
	)
	
	var f *os.File
	
	if _, err := os.Stat(pngfile); err == nil {
		os.Remove(pngfile)
	}
	
	f, err := os.Create(pngfile)
	if err != nil {
		fmt.Printf("Create file failed%v \n", err)
		os.Exit(1)
	}
	
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py) /height*(ymax - ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax - xmin) + xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(f, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15
	
	var v complex128
	
	for n := uint8(0); n < contrast; n++ {
		v = v * v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	
	return color.Black
}
