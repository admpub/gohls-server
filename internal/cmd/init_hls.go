package cmd

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/admpub/gohls-server/internal/hls"
	homedir "github.com/mitchellh/go-homedir"
	log "github.com/sirupsen/logrus"
)

func init_hls() {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.InfoLevel)
	if _, err := strconv.ParseBool(os.Getenv("DEBUG")); err == nil {
		log.SetLevel(log.DebugLevel)
	}

	// Find ffmpeg
	ffmpeg, err := exec.LookPath("ffmpeg")
	if err != nil {
		log.Fatal("ffmpeg could not be found in your path", err)
	}

	// Find ffprobe
	ffprobe, err := exec.LookPath("ffprobe")
	if err != nil {
		log.Fatal("ffprobe could not be found in your path", err)
	}

	homeDir, err := homedir.Dir()
	if err != nil {
		log.Fatal("Could not determine home directory", err)
	}

	// Configure HLS module
	hls.FFMPEGPath = ffmpeg
	hls.FFProbePath = ffprobe
	hls.HomeDir = filepath.Join(homeDir, ".gohls")
}
