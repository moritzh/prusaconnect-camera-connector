package main

import (
	"github.com/moritzh/prusa_webcam_thing/internal/camera"
)

func main() {

	cameraManager := camera.LoadConfiguration()

	cameraManager.UploadAll()

}
