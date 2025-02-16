package camera

import (
	"fmt"
	"time"

	"gopkg.in/ini.v1"
)

type CameraManager struct {
	cameras        []CameraConnection
	updateInterval int
}

func LoadConfiguration() *CameraManager {
	inidata, err := ini.Load("config.ini")

	if err != nil {
		fmt.Printf("Unable to load config.ini, does it exist?")
		return &CameraManager{}
	}

	updateInterval := inidata.Section("").Key("interval").MustInt(30)

	cameras := make([]CameraConnection, 0)

	for _, section := range inidata.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}
		token := section.Key("token").Value()
		fingerprint := section.Key("fingerprint").Value()
		cameraDevice := section.Key("device").Value()
		strategy := "ffmpeg"
		fmt.Println(section.Name())

		cameras = append(cameras, *NewCameraConnection(section.Name(), token, fingerprint, cameraDevice, strategy))

	}

	return &CameraManager{cameras: cameras, updateInterval: updateInterval}
}

func (c *CameraManager) StartUploading() {
	for {
		count := 0
		for _, camera := range c.cameras {
			success := camera.Upload()
			if success {
				count = count + 1
			}

		}
		fmt.Printf("Uploaded %d images, will now wait %d seconds\n", count, c.updateInterval)
		time.Sleep(time.Duration(c.updateInterval) * time.Second)
	}

}
