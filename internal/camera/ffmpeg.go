package camera

import (
	"os"
	"os/exec"
)

// ffmpeg -f video4linux2 -s 640x480 -i /dev/video0 -ss 0:0:2 -frames 1 /tmp/out.jpg

func FFMpegCaptureImage(device string) (*os.File, error) {
	tmpFile, _ := os.CreateTemp("/tmp", "*.jpg")
	command := "ffmpeg"

	var cmd *exec.Cmd = nil

	// I can't get this to work, since Mac is not showing the permission dialog to actually allow screen capturing. There's probably some ways to make this work, but since this will run on a raspberry, i don't care too much. Mostly for testing.
	// if runtime.GOOS == "darwin" {
	// 	cmd = exec.Command(command, "-f", "avfoundation", "-i", "\"FaceTime HD Camera\"", "-framerate", "30", "-s", "1024x786", "-ss", "0:0:2", "-frames", "1", tmpFile.Name())
	// } else {
	cmd = exec.Command(command, "-y", "-f", "video4linux2", "-s", "1024x786", "-i", device, "-ss", "0:0:2", "-frames", "1", tmpFile.Name())
	//}

	cmd.Stdout = os.Stdout // or any other io.Writer
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, err
	}

	return tmpFile, nil
}
