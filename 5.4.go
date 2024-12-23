//задание 4

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"sync"
)

func applyKernel(img draw.Image, x, y int, kernel [][]float64) color.RGBA64 {
	bounds := img.Bounds()
	kernelSize := len(kernel)
	offset := kernelSize / 2
	var r, g, b float64

	for ky := 0; ky < kernelSize; ky++ {
		for kx := 0; kx < kernelSize; kx++ {
			ix := x + kx - offset
			iy := y + ky - offset
			if ix >= bounds.Min.X && ix < bounds.Max.X && iy >= bounds.Min.Y && iy < bounds.Max.Y {
				pixel := img.At(ix, iy).(color.RGBA64)
				weight := kernel[ky][kx]
				r += float64(pixel.R) * weight
				g += float64(pixel.G) * weight
				b += float64(pixel.B) * weight
			}
		}
	}

	return color.RGBA64{
		R: uint16(r),
		G: uint16(g),
		B: uint16(b),
		A: 65535,
	}
}

func filterWithKernel(img draw.Image, kernel [][]float64, y int, wg *sync.WaitGroup) {
	defer wg.Done()
	bounds := img.Bounds()
	copyImg := image.NewRGBA64(bounds)
	draw.Draw(copyImg, bounds, img, bounds.Min, draw.Src)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		newColor := applyKernel(copyImg, x, y, kernel)
		img.Set(x, y, newColor)
	}
}

func main() {
	// Открываем файл с изображением
	file, err := os.Open("input.png")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Декодируем изображение
	img, _, err := image.Decode(file)
	if err != nil {
		fmt.Println("Error decoding image:", err)
		return
	}

	// Преобразуем изображение в редактируемый формат
	drawImg, ok := img.(draw.Image)
	if !ok {
		fmt.Println("Image conversion failed")
		return
	}

	// Определяем ядро свертки (например, размытие по Гауссу)
	kernel := [][]float64{
		{1 / 16.0, 2 / 16.0, 1 / 16.0},
		{2 / 16.0, 4 / 16.0, 2 / 16.0},
		{1 / 16.0, 2 / 16.0, 1 / 16.0},
	}

	// Применяем фильтр параллельно
	var wg sync.WaitGroup
	bounds := drawImg.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		wg.Add(1)
		go filterWithKernel(drawImg, kernel, y, &wg)
	}
	wg.Wait()

	// Сохраняем обработанное изображение
	output, err := os.Create("output_kernel.png")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer output.Close()

	err = png.Encode(output, drawImg)
	if err != nil {
		fmt.Println("Error saving image:", err)
	}
}
