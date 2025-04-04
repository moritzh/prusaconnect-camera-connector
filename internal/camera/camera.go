package camera

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type CameraConnection struct {
	Nameconst       string
	Token           string // the token as indicated from the prusa website.
	Fingerprint     string // this is locally sourced and created.
	LocalCameraName string // something like /dev/video0
	StrategyName    string // very optional, only ffmpeg for now. but who knows.
}

func NewCameraConnection(name string, token string, fingerprint string, localCameraName string, strategyName string) *CameraConnection {
	return &CameraConnection{
		name, token, fingerprint, localCameraName, strategyName,
	}
}

func (c *CameraConnection) Upload() {
	imageFile, error := FFMpegCaptureImage(c.LocalCameraName)

	if error != nil {
		log.Print(error)
	} else {
		c.uploadSingleFileImage(imageFile)
		imageFile.Close()
		os.Remove(imageFile.Name())
	}

}

func (c *CameraConnection) uploadSingleFileImage(file *os.File) {
	info, _ := file.Stat()
	request, _ := http.NewRequest("PUT", "https://connect.prusa3d.com/c/snapshot", file)
	request.ContentLength = info.Size()

	request.Header.Set("token", c.Token)
	request.Header.Set("fingerprint", c.Fingerprint)
	request.Header.Set("Content-Type", "image/jpeg")

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		log.Printf("Upload failed with error: %w", err)
		return
	} else {
		defer res.Body.Close()

		fmt.Printf("Upload Done, Status Code %d\n", res.StatusCode)
	}
}
