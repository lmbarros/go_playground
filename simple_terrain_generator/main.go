package main

import (
	"errors"
	"image"
	"image/color"
	"image/png"
	"math"
	"net/http"
	"os"

	"github.com/lmbarros/sbxs_go_noise/fractalnoise"
	"github.com/lmbarros/sbxs_go_noise/opensimplex"
	"github.com/lmbarros/sbxs_go_rand/randutil"
)

const (
	// mapWidth is the final image width, in pixels.
	mapWidth = 1024

	// mapHeight is the final image height, in pixels.
	mapHeight = 768

	// seaLevel is the sea level (heights range from zero to one).
	seaLevel = 0.45
)

var (
	// deepWaterColor is the color of the deepest waters
	deepWaterColor = color.NRGBA{10, 20, 125, 255}

	// shallowWaterColor is the color of the shallowest waters
	shallowWaterColor = color.NRGBA{30, 50, 200, 255}
)

// mixColors interpolates colors c1 and c2 with a given ratio.
func mixColors(c1, c2 *color.NRGBA, ratio float64) (color.NRGBA, error) {
	if ratio < 0.0 || ratio > 1.0 {
		return color.NRGBA{}, errors.New("Invalid ration for mixing colors")
	}

	w1 := ratio
	w2 := 1.0 - ratio

	return color.NRGBA{
		byte(float64(c1.R)*w1 + float64(c2.R)*w2),
		byte(float64(c1.G)*w1 + float64(c2.G)*w2),
		byte(float64(c1.B)*w1 + float64(c2.B)*w2),
		byte(float64(c1.A)*w1 + float64(c2.A)*w2)}, nil
}

// hc is a height and color pair
type hc struct {
	h float64
	c color.NRGBA
}

// heightToColor converts a height (from zero to one) to a nice color.
func heightToColor(height float64) color.NRGBA {

	// Heights and colors
	hcs := []hc{
		{0.0, shallowWaterColor},               // just to have a smoother shores
		{0.05, color.NRGBA{10, 200, 50, 255}},  // light green
		{0.35, color.NRGBA{8, 150, 40, 255}},   // dark green
		{0.5, color.NRGBA{255, 240, 25, 255}},  // yellow
		{0.65, color.NRGBA{133, 94, 43, 255}},  // brown
		{0.85, color.NRGBA{166, 120, 60, 255}}, // still brown, just lighter
		{1.0, color.NRGBA{220, 220, 220, 255}}, // light gray
	}

	// Are we under water?
	if height < seaLevel {
		r := height / seaLevel
		seaColor, _ := mixColors(&shallowWaterColor, &deepWaterColor, r)
		return seaColor
	}

	// Are we above sea level (ASL)?
	aslHeight := (height - seaLevel) / (1.0 - seaLevel)

	var largerIndex int

	for i, hc := range hcs {
		if hc.h == aslHeight {
			return hc.c
		} else if hc.h > aslHeight {
			largerIndex = i
			break
		}
	}

	sh := hcs[largerIndex-1].h // small height
	lh := hcs[largerIndex].h   // large height
	r := (aslHeight - sh) / (lh - sh)

	landColor, _ := mixColors(&hcs[largerIndex].c, &hcs[largerIndex-1].c, r)
	return landColor
}

// createMap returns a new random map, as an image.
func createMap(params fractalnoise.Params, seed int64) *image.NRGBA {
	// Create a height map with floating point numbers
	var heightMap [mapWidth][mapHeight]float64

	noiseGen := fractalnoise.New2D(opensimplex.NewWithSeed(seed), params)

	minHeight := math.MaxFloat64
	maxHeight := -math.MaxFloat64

	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			height := noiseGen.Noise2D(float64(x), float64(y))
			if height > maxHeight {
				maxHeight = height
			}

			if height < minHeight {
				minHeight = height
			}

			heightMap[x][y] = height
		}
	}

	// Normalize the height map
	deltaHeight := maxHeight - minHeight
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			heightMap[x][y] = (heightMap[x][y] - minHeight) / deltaHeight
		}
	}

	minHeight = math.MaxFloat64
	maxHeight = -math.MaxFloat64
	for y := 0; y < mapHeight; y++ {
		for x := 0; x < mapWidth; x++ {
			if heightMap[x][y] > maxHeight {
				maxHeight = heightMap[x][y]
			}
			if heightMap[x][y] < minHeight {
				minHeight = heightMap[x][y]
			}
		}
	}

	// Create the image
	image := image.NewNRGBA(image.Rect(0, 0, mapWidth, mapHeight))
	b := image.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			image.Set(x, y, heightToColor(heightMap[x][y]))
		}
	}

	// There we are
	return image
}

// handleEverything handles all HTTP requests.
func handleEverything(w http.ResponseWriter, r *http.Request) {
	// Create the map
	params := fractalnoise.Params{
		Layers:     9,
		Frequency:  0.006,
		Lacunarity: 1.55,
		Gain:       0.75,
	}

	image := createMap(params, randutil.GoodSeed())

	// Save the map to a PNG file, just because I want to keep it around
	{
		pngFile, err := os.Create("terrain.png")

		if err != nil {
			panic(err)
		}

		defer pngFile.Close()

		png.Encode(pngFile, image)
	}

	// Return the image to the caller
	png.Encode(w, image)
}

// main is the entry point.
func main() {
	http.HandleFunc("/", handleEverything)
	http.ListenAndServe(":8080", nil)
}
