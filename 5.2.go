//2 задание

package main

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"time"
)

func filter(img draw.Image) {
	bounds := img.Bounds()
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			original := img.At(x, y).(color.RGBA64)
			gray := uint16((uint32(original.R) + uint32(original.G) + uint32(original.B)) / 3)
			img.Set(x, y, color.RGBA64{R: gray, G: gray, B: gray, A: original.A})
		}
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

	// Применяем фильтр
	start := time.Now()
	filter(drawImg)
	fmt.Println("Processing time:", time.Since(start))

	// Сохраняем обработанное изображение
	output, err := os.Create("output.png")
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
