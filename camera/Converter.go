package camera

import (
	"github.com/babell00/toc_camera/configuration"
	"image"
	"os"
	"log"
)

func ConvertConfigCameraToCamera(config configuration.Config) []Camera {
	var cameras []Camera
	for _, v := range config.Cameras {
		img := openImage(v.ErrorImage, config)
		c := Camera{Name: v.Name, Url: v.Url, Path: v.Path, Image: img}
		cameras = append(cameras, c)
	}
	return cameras
}

func openImage(fileName string, config configuration.Config) image.Image {
	f, err := os.Open(fileName)
	defer  f.Close()
	if err != nil {
		log.Println("Cannot open image:", fileName)
		f, _ = os.Open(config.DefaultErrorImage)
	}
	img, _, err := image.Decode(f)
	if err != nil {
		log.Println("Cannot load file:", fileName)
	}
	return img
}
