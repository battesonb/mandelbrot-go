package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/cmplx"
	"os"
	"sync"
)

const (
	WIDTH     = 1280 * 2
	HEIGHT    = 720 * 2
	MAX_STEPS = 3000
	MAX_ZOOM  = 11.0
)

type GradientColor struct {
	Stop float64
	R    uint8
	G    uint8
	B    uint8
}

func main() {
	for i := range 100 {
		ratio := float64(i) / 100.0
		zoom := MAX_ZOOM * ratio
		mandelbrot(math.Pow(2.0, zoom), i)
	}
}

func mandelbrot(zoom float64, frame int) {
	image := image.NewRGBA(image.Rect(0, 0, WIDTH, HEIGHT))
	var wg sync.WaitGroup
	for y := range HEIGHT {
		wg.Add(1)
		go func(y int) {
			defer wg.Done()
			for x := range WIDTH {
				centerX := 0.577
				centerY := 0.357
				c := complex(
					scale(scale(float64(x)/float64(WIDTH), centerX-(1/zoom)/2, centerX+(1/zoom)/2), -2.3, 0.5),
					scale(scale(float64(y)/float64(HEIGHT), centerY-(1/zoom)/2, centerY+(1/zoom)/2), -1.14, 1.14),
				)
				z := complex(0, 0)
				var steps int
				for steps = range MAX_STEPS {
					if cmplx.Abs(z) > 2 {
						break
					}
					z = z*z + c
				}

				intensity := (float64(steps) / float64(MAX_STEPS-1))
				image.SetRGBA(x, y, colorForIntensity(intensity))
			}
		}(y)
	}

	wg.Wait()

	err := saveImage(fmt.Sprintf("./out/mandelbrot-%03d.png", frame), image)
	if err != nil {
		panic(err)
	}
}

func colorForIntensity(intensity float64) color.RGBA {
	for i, a := range GRADIENT {
		if a.Stop > intensity {
			if i == 0 {
				return color.RGBA{
					R: a.R,
					G: a.G,
					B: a.B,
					A: 255,
				}
			}

			b := GRADIENT[i-1]
			intensity = (intensity - b.Stop) / (a.Stop - b.Stop)
			return color.RGBA{
				R: uint8(intensity*float64(a.R)) + uint8((1-intensity)*float64(b.R)),
				G: uint8(intensity*float64(a.G)) + uint8((1-intensity)*float64(b.G)),
				B: uint8(intensity*float64(a.B)) + uint8((1-intensity)*float64(b.B)),
				A: 255,
			}
		}
	}
	c := GRADIENT[len(GRADIENT)-1]
	return color.RGBA{
		R: c.R,
		G: c.G,
		B: c.B,
		A: 255,
	}
}

func scale(ratio float64, min float64, max float64) float64 {
	return (max-min)*math.Max(0, math.Min(1, ratio)) + min
}

func saveImage(path string, image image.Image) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	err = png.Encode(file, image)
	if err != nil {
		return err
	}

	err = file.Sync()
	if err != nil {
		return err
	}

	return nil
}
