package camera

import (
	"errors"
	"fmt"
	"os"
	"time"

	"gopkg.in/ini.v1"
)

type CameraManager struct {
	cameras        []CameraConnection
	updateInterval int
}

func loadConfigurationFile() (*ini.File, error) {

	paths := []string{"/etc/pccc.config", "./pccc.config"}

	for _, path := range paths {
		data, _ := ini.Load(path)
		if data != nil {
			return data, nil
		}
	}

	return nil, errors.New("No config file found in expected location")
}

func LoadConfiguration() *CameraManager {
	inidata, err := loadConfigurationFile()

	if err != nil {
		fmt.Printf("Can't find config.ini – Please consult the Readme at https://github.com/moritzh/prusa-webcam-uploader to get started.\n\n")
		os.Exit(1)
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
		cameras = append(cameras, *NewCameraConnection(section.Name(), token, fingerprint, cameraDevice, strategy))
	}

	if len(cameras) > 0 {
		fmt.Printf("Found %d cameras\n", len(cameras))
	} else {
		fmt.Println("No cameras found, ciao!")
		os.Exit(1)

	}

	return &CameraManager{cameras: cameras, updateInterval: updateInterval}
}

func (c *CameraManager) StartUploading() {
	for {

		for _, camera := range c.cameras {
			camera.Upload()
		}

		fmt.Printf("Done with image uploads, waiting %d seconds.\n", c.updateInterval)
		time.Sleep(time.Duration(c.updateInterval) * time.Second)
	}

}
