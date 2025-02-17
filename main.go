package main

import (
	"sync"

	"github.com/moritzh/prusaconnect-camera-connector/internal/camera"
)

func main() {

	cameraManager := camera.LoadConfiguration()
	var wg sync.WaitGroup
	wg.Add(1)
	go cameraManager.StartUploading()
	wg.Wait()
}
