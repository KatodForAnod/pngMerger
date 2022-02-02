package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type LoadedConf struct {
	Images []ImgSpecs
}

type ImgSpecs struct {
	Filename    string `json:"filename"`
	ColorOld    string `json:"color_old"`
	ColorNew    string `json:"color_new"`
	PaintPointX int    `json:"paint_point_x"`
	PaintPointY int    `json:"paint_point_y"`
}

const fileName = "config.conf"

func LoadConfig() (loadedConf LoadedConf, err error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Println(err)
		return LoadedConf{}, err
	}

	err = json.Unmarshal(data, &loadedConf)
	return
}
