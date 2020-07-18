package generator

import (
	"image"
	"image/color"
	"image/draw"
	"math/rand"
	"time"
)

func GenerateIcon(hash []byte, size int) *image.RGBA {
	r, g, b := pickColor()
	drawColor := color.RGBA{r, g, b, 255}
	bgColor := color.RGBA{255, 255, 255, 255}

	imageRect := image.Rect(0, 0, size, size)
	result := image.NewRGBA(imageRect)
	draw.Draw(result, imageRect, &image.Uniform{bgColor}, image.ZP, draw.Src)

	grid := buildGrid(hash)
	pixelMap := buildPixelMap(grid, size/5)
	for _, rect := range pixelMap {
		draw.Draw(result, rect, &image.Uniform{drawColor}, image.ZP, draw.Src)
	}

	return result
}

func pickColor() (red, green, blue uint8) {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return uint8(r.Intn(256)), uint8(r.Intn(256)), uint8(r.Intn(256))
}

func buildPixelMap(grid []byte, size int) []image.Rectangle {
	var result []image.Rectangle

	for index, item := range grid {
		if item%2 == 1 {
			x := (index % 5) * size
			y := (index / 5) * size

			result = append(result, image.Rect(x, y, x+size, y+size))
		}
	}

	return result
}

func buildGrid(data []byte) []byte {
	var result []byte

	for i := 0; i < len(data); i += 3 {
		end := i + 3
		if end > len(data) {
			break
		}

		row := append([]byte(nil), data[i:end]...)
		for j := end - 2; j >= i; j-- {
			row = append(row, data[j])
		}

		result = append(result, row...)
	}

	return result
}
