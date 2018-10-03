package main

import (
	"flag"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/admpub/gohls-server/hls"
)

func setVideoDir(f *flag.FlagSet) string {
	if f.NArg() > 0 {
		videoDir := f.Arg(0)
		hls.HomeDir = filepath.Join(videoDir, ".gohls")
		return videoDir
	}
	log.Fatalf("Path to videos not specified")
	return ""
}
