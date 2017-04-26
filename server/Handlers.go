package server

import (
	"net/http"
	"github.com/babell00/toc_camera/camera"
	"sort"
	"encoding/json"
	"image/jpeg"
	"github.com/gorilla/mux"
	"log"
	"mime/multipart"
	"net/textproto"
	"bytes"
	"fmt"
	"sync"
	"image"
)

var _service *camera.CameraService

func InitHandlers(service *camera.CameraService) {
	_service = service

}

func CameraIndex(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving", r.URL)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	cameras := getService().GetAll()

	sort.Slice(cameras, func(a, b int) bool {
		return cameras[a].Name < cameras[b].Name
	})

	json.NewEncoder(w).Encode(&cameras)
}

func CameraImage(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving", r.URL)

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "image/jpeg")

	vars := mux.Vars(r)
	path := vars["camera_path"]

	camera := getService().GetByPath(path)
	if camera.Image == nil {
		return
	}

	err := jpeg.Encode(w, camera.Image, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Experimental
func CameraStream(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving streaming", r.URL)

	m := multipart.NewWriter(w)
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary="+ m.Boundary())
	w.Header().Set("Connection", "close")

	vars := mux.Vars(r)
	path := vars["camera_path"]
	cam := getService().GetByPath(path)
	header := textproto.MIMEHeader{}

	var buf bytes.Buffer
	decoder ,_ := camera.Decoder(cam.MJpegUrl)
	var mutex sync.RWMutex
	img, _:= camera.ReadJpeg(cam.MJpegUrl)
	go func() {
		for {
			var tmp image.Image
			tmp, err := decoder.Decode()
			if err != nil {
				break
			}
			mutex.Lock()
			img = tmp
			mutex.Unlock()
		}
	}()

	for {
		mutex.RLock()
		buf.Reset()
		err := jpeg.Encode(&buf, img, nil)
		mutex.RUnlock()
		if err != nil {
			break
		}
		header.Set("Content-Type", "image/jpeg")
		header.Set("Content-Length", fmt.Sprint(buf.Len()))
		mw, err := m.CreatePart(header)
		if err != nil {
			break
		}
		mw.Write(buf.Bytes())
		if flusher, ok := mw.(http.Flusher); ok {
			flusher.Flush()
		}
	}

}


func getService() *camera.CameraService {
	if _service == nil {
		panic("Handlers need to be initialize. InitHandlers(ser *camera.CameraService)")
	}
	return _service
}
