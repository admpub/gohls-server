package utils

import (
	"github.com/admpub/gohls-server/internal/hls"
)

var (
	ConvertToMP4  = hls.ConvertToMP4
	IsUnsupported = hls.IsUnsupported
)

func HomeDir() string {
	return hls.HomeDir
}
func FFProbePath() string {
	return hls.FFProbePath
}
func FFMPEGPath() string {
	return hls.FFMPEGPath
}
func ComSkipINI() string {
	return hls.ComSkipINI
}
func ComSkipPath() string {
	return hls.ComSkipPath
}

func SetHomeDir(homeDir string) {
	hls.HomeDir = homeDir
}

func SetFFProbePath(ffProbePath string) {
	hls.FFProbePath = ffProbePath
}

func SetFFMPEGPath(ffMPEGPath string) {
	hls.FFMPEGPath = ffMPEGPath
}

func SetComSkipINI(iniFile string) {
	hls.ComSkipINI = iniFile
}

func SetComSkipPath(fpath string) {
	hls.ComSkipPath = fpath
}
