package camera

import (
	"github.com/babell00/toc_camera/configuration"
	"image"
	"os"
	"log"
	"fmt"
)

func ConvertConfigCameraToCamera(config configuration.Config) []Camera {
	var cameras []Camera
	for _, v := range config.Cameras {
		img := openImage(v.ErrorImage, config)
		publicImageUrl := buildPublicImageUrl(config.Server.PublicAddress, v.UrlPath)
		publicStreamUrl := buildPublicStreamUrl(config.Server.PublicAddress, v.UrlPath)
		c := Camera{Name: v.Name, MJpegUrl: v.MJpegUrl, UrlPath: v.UrlPath, PublicImageUrl: publicImageUrl, PublicStreamUrl: publicStreamUrl, Image: img}
		cameras = append(cameras, c)
	}
	return cameras
}

func buildPublicImageUrl(publicUrl string, path string) string {
	image := "image"
	return fmt.Sprintf("http://%v/%v/%v", publicUrl, image, path)
}

func buildPublicStreamUrl(publicUrl string, path string) string {
	image := "stream"
	return fmt.Sprintf("http://%v/%v/%v", publicUrl, image, path)
}

func openImage(fileName string, config configuration.Config) image.Image {
	f, err := os.Open(fileName)
	defer f.Close()
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
