package main

import (
	"image"
	"image/png"
	"log"
	"os"
	"pngMerger/config"
	"pngMerger/imager"
)

func main() {
	loadConfig, err := config.LoadConfig()
	if err != nil {
		return
	}

	images := make([]image.Image, 0, 0)
	for _, specs := range loadConfig.Images {
		pngImage, err := os.Open(specs.Filename)
		if err != nil {
			panic(err.Error())
		}
		src, _, err := image.Decode(pngImage)
		if err != nil {
			panic(err.Error())
		}
		pngImage.Close()

		if specs.ColorOld != "" && specs.ColorNew != "" {
			newImg, err := imager.ReplaceHexColors(src, specs.ColorOld, specs.ColorNew)
			if err != nil {
				log.Println(err)
			} else {
				images = append(images, newImg)
				continue
			}
		}

		images = append(images, src)
	}

	mergedImg, err := imager.MergePng(images)
	if err != nil {
		panic(err.Error())
	}

	outfile, err := os.Create("newPng.png")
	if err != nil {
		panic(err.Error())
	}
	defer outfile.Close()

	png.Encode(outfile, mergedImg)
}
