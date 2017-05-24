package server

import (
	"net/http"
	"github.com/babell00/tCamera/camera"
	"sort"
	"encoding/json"
	"image/jpeg"
	"github.com/gorilla/mux"
	"log"
	"net/textproto"
	"bytes"
	"mime/multipart"
	"image"
	"sync"
	"fmt"
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
	//log.Println("Serving", r.URL)

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
		log.Println(err)
	}
}

//Experimental
func CameraStream(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving streaming", r.URL)

	m := multipart.NewWriter(w)
	w.Header().Set("Content-Type", "multipart/x-mixed-replace; boundary="+m.Boundary())
	w.Header().Set("Connection", "close")
	vars := mux.Vars(r)
	path := vars["camera_path"]
	cam := getService().GetByPath(path)
	header := textproto.MIMEHeader{}

	var buf bytes.Buffer
	decoder, _ := camera.Decoder(cam.MJpegUrl)

	var img image.Image
	var mux sync.Mutex
	for {
		mux.Lock()
		tmp, err := decoder.Decode()
		if err != nil {
			log.Println(err)
			break
		}
		img = tmp
		mux.Unlock()

		buf.Reset()
		mux.Lock()
		err = jpeg.Encode(&buf, img, nil)

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
		mux.Unlock()
	}

}

func getService() *camera.CameraService {
	if _service == nil {
		panic("Handlers need to be initialize. InitHandlers(ser *camera.CameraService)")
	}
	return _service
}
