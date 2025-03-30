package camera

import (
	"context"
	"os"
	"os/exec"
	"time"
)

// ffmpeg -f video4linux2 -s 640x480 -i /dev/video0 -ss 0:0:2 -frames 1 /tmp/out.jpg

func FFMpegCaptureImage(device string) (*os.File, error) {
	tmpFile, _ := os.CreateTemp("/tmp", "*.jpg")
	command := "ffmpeg"

	var cmd *exec.Cmd = nil

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cmd = exec.CommandContext(ctx, command, "-threads", "1", "-y", "-f", "video4linux2", "-s", "1024x786", "-i", device, "-ss", "0:0:2", "-frames", "1", tmpFile.Name())

	// uncomment for debugging, otherwise a little noisy?
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
		return nil, err
	}

	return tmpFile, nil
}
