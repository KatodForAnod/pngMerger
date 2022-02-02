package imager

import (
	"errors"
	"image"
	"image/color"
	"log"
	"strconv"
)

func MergePng(images []image.Image) (*image.RGBA, error) {
	h, w, err := FindMinSizeHW(images)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	outputBounds := image.Rectangle{
		Min: image.Point{},
		Max: image.Point{
			X: w,
			Y: h,
		},
	}

	outputImage := image.NewRGBA(outputBounds)
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			for _, img := range images {
				oldColor := img.At(x, y)
				if _, _, _, alpha := oldColor.RGBA(); alpha != 0 {
					continue
				}
				outputImage.Set(x, y, oldColor)
			}
		}
	}

	return outputImage, nil
}

func FindMinSizeHW(images []image.Image) (h, w int, err error) {
	if len(images) < 1 {
		return 0, 0, errors.New("empty array")
	}

	h = images[0].Bounds().Max.Y
	w = images[0].Bounds().Max.X

	for _, img := range images {
		bounds := img.Bounds()

		if bounds.Max.X <= w {
			w = bounds.Max.X
		}
		if bounds.Max.Y <= h {
			h = bounds.Max.Y
		}
	}

	return h, w, nil
}

func ReplaceHexColors(img image.Image, colorHexBefore, colorHexAfter string) (*image.RGBA, error) {
	bounds := img.Bounds()
	outputImage := image.NewRGBA(bounds)

	colorBefore, err := Hex2Color(colorHexBefore)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	colorAfter, err := Hex2Color(colorHexAfter)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			currColor := img.At(x, y)
			if _, _, _, alpha := currColor.RGBA(); alpha != 0 {
				continue
			}
			if currColor == colorBefore {
				outputImage.Set(x, y, colorAfter)
			} else {
				outputImage.Set(x, y, currColor)
			}
		}
	}

	return outputImage, nil
}

func Hex2RGB(hex string) (uint8, uint8, uint8, error) {

	values, err := strconv.ParseUint(hex, 16, 32)
	if err != nil {
		return 0, 0, 0, err
	}

	Red := uint8(values >> 16)
	Green := uint8((values >> 8) & 0xFF)
	Blue := uint8(values & 0xFF)

	return Red, Green, Blue, nil
}

func Hex2Color(hex string) (color.Color, error) {
	r, g, b, err := Hex2RGB(hex)
	if err != nil {
		log.Println(err)
		return color.RGBA{}, err
	}

	return color.RGBA{r, g, b, 0xff}, nil
}
