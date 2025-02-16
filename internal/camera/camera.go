package camera

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2/log"
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
		log.Error(error)
	}

	//	if imageFile != nil {
	c.uploadSingleFileImage(imageFile)
	defer imageFile.Close()
	fmt.Println("Uploaded a File!")
	//}

}

func (c *CameraConnection) uploadSingleFileImage(file *os.File) {

	client := http.Client{}

	testFile, err := os.Open("image.jpg")

	if err != nil {
		fmt.Print(err)
	}

	//request, _ := http.NewRequest("PUT", "http://192.168.178.21:8080", testFile)
	info, _ := testFile.Stat()
	request, _ := http.NewRequest("PUT", "https://connect.prusa3d.com/c/snapshot", testFile)
	request.ContentLength = info.Size()

	request.Header.Set("token", c.Token)
	request.Header.Set("fingerprint", c.Fingerprint)
	request.Header.Set("Content-Type", "image/jpeg")

	fmt.Println(c.Token)
	fmt.Println(c.Fingerprint)

	res, err := client.Do(request)

	if err != nil {
		log.Error(err)

	}

	response := make([]byte, res.ContentLength)
	res.Body.Read(response)
	fmt.Println(string(response))

	fmt.Printf("Upload Done, Status Code %d", res.StatusCode)
}
