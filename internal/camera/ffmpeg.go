package camera

import (
	"os"
	"os/exec"
)

// ffmpeg -f video4linux2 -s 640x480 -i /dev/video0 -ss 0:0:2 -frames 1 /tmp/out.jpg

func FFMpegCaptureImage(device string) (*os.File, error) {
	tmpFile, _ := os.CreateTemp("/tmp", "*.jpg")
	command := "ffmpeg"

	cmd := exec.Command(command, "-f video4linux2", "-s 1024x786", "-i", device, "-ss 0:0:2", "-frames", "1", tmpFile.Name())

	if err := cmd.Run(); err != nil {
		tmpFile.Close()
		return nil, err
	}

	return tmpFile, nil
}
