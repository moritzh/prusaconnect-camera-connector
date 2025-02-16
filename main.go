package main

import (
	"sync"

	"github.com/moritzh/prusa_webcam_thing/internal/camera"
)

func main() {

	cameraManager := camera.LoadConfiguration()
	var wg sync.WaitGroup
	wg.Add(1)
	go cameraManager.StartUploading()
	wg.Wait()
}
