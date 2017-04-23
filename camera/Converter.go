package camera

import (
	"github.com/babell00/toc_camera/configuration"
)

func ConvertConfigCameraToCamera(configCameras []configuration.Camera) []Camera {
	var cameras []Camera
	for _, v := range configCameras {
		c := Camera{Name: v.Name, Url: v.Url, Path: v.Path}
		cameras = append(cameras, c)
	}
	return cameras
}
