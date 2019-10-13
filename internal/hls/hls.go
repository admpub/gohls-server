package hls

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"

	"github.com/admpub/errors"
)

var (
	HomeDir     = ".gohls"
	FFProbePath = "ffprobe"
	FFMPEGPath  = "ffmpeg"

	//ComSkipINI comskip's ini file path
	ComSkipINI  = ""
	ComSkipPath = "comskip"
)

const (
	cacheDirName     = "cache"
	hlsSegmentLength = 5.0 // Seconds
)

func ClearCache() error {
	var cacheDir = filepath.Join(HomeDir, cacheDirName)
	return os.RemoveAll(cacheDir)
}

func ConvertToMP4(videoFile string, outputFile string) error {
	if _, err := exec.LookPath(FFMPEGPath); err != nil {
		return errors.New("Cannot find " + FFMPEGPath + " executable in path")
	}
	size := 1
	if _, err := exec.LookPath(ComSkipPath); err != nil {
		fmt.Println("Cannot find " + ComSkipPath + " executable in path")
	} else {
		size++
	}
	ch := make(chan error, size)
	go func() {
		args := []string{"-fflags", "+discardcorrupt", "-i", videoFile, "-c:a", "copy", "-bsf:a", "aac_adtstoasc", "-c:v", "h264_omx", "-b:v", "5000k", "-y", outputFile}
		_, err := execute(FFMPEGPath, args)
		ch <- err
	}()
	if size > 1 {
		go func() {
			args := []string{"-d", "255", "--ini=" + ComSkipINI, "--threads=" + strconv.Itoa(runtime.NumCPU()), "--hwassist", "-t", outputFile}
			_, err := execute(ComSkipPath, args)
			ch <- err
		}()
	}
	var err error
	for i := 0; i < size; i++ {
		recvErr := <-ch
		if recvErr != nil {
			err = recvErr
		}
	}
	if err != nil {
		return errors.WithMessage(err, "Some error occurred during encoding or detecting commercials")
	}
	return err
}
