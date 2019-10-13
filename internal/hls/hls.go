package hls

import (
	"database/sql"
	"fmt"
	"log"
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

	supportedFFProbe sql.NullBool
	supportedFFMPEG  sql.NullBool
	supportedComSkip sql.NullBool
)

func Reset() {
	supportedFFProbe.Valid = false
	supportedFFMPEG.Valid = false
	supportedComSkip.Valid = false
}

const (
	cacheDirName     = "cache"
	hlsSegmentLength = 5.0 // Seconds
)

func ClearCache() error {
	var cacheDir = filepath.Join(HomeDir, cacheDirName)
	return os.RemoveAll(cacheDir)
}

func IsUnsupported(err error) bool {
	return errors.Cause(err) == ErrUnsupported
}

func IsSupportedFFProbe() bool {
	if supportedFFProbe.Valid {
		return supportedFFProbe.Bool
	}
	supportedFFProbe.Valid = true
	supportedFFProbe.Bool = false
	if _, err := exec.LookPath(FFProbePath); err != nil {
		return false
	}
	supportedFFProbe.Bool = true
	return true
}

func IsSupportedFFMPEG() bool {
	if supportedFFMPEG.Valid {
		return supportedFFMPEG.Bool
	}
	supportedFFMPEG.Valid = true
	supportedFFMPEG.Bool = false
	if _, err := exec.LookPath(FFMPEGPath); err != nil {
		return false
	}
	supportedFFMPEG.Bool = true
	return true
}

func IsSupportedComSkip() bool {
	if supportedComSkip.Valid {
		return supportedComSkip.Bool
	}
	supportedComSkip.Valid = true
	supportedComSkip.Bool = false
	if len(ComSkipINI) == 0 {
		return false
	}
	if _, err := os.Stat(ComSkipINI); err != nil {
		log.Println(err)
		return false
	}
	if _, err := exec.LookPath(ComSkipPath); err != nil {
		return false
	}
	supportedComSkip.Bool = true
	return true
}

func ConvertToMP4(videoFile string, outputFile string) error {
	if !IsSupportedFFMPEG() {
		return errors.WithMessage(ErrUnsupported, "Cannot find "+FFMPEGPath+" executable in path")
	}
	size := 1
	if IsSupportedComSkip() {
		size++
	} else {
		fmt.Println("Cannot find " + ComSkipPath + " executable in path")
	}
	ch := make(chan error, size)
	go func() {
		args := []string{"-i", videoFile, "-acodec", "copy", "-vcodec", "copy", outputFile}
		//ffmpeg -i index.ts -acodec copy -vcodec copy index.mp4
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
