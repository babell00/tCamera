package camera

import (
	"github.com/babell00/toc_camera/configuration"
	"github.com/twinj/uuid"
)

func ConvertConfigCameraToCamera(configCameras []configuration.Camera) []Camera {
	var cameras []Camera
	for _, v := range configCameras {
		u := uuid.NewV4().String()
		c := Camera{Id: u, Name: v.Name, Url: v.Url, Path: v.Path}
		cameras = append(cameras, c)
	}
	return cameras
}
