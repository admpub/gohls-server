package utils

import (
	"testing"
)

func TestConvert(t *testing.T) {
	return
	err := ConvertToMP4(`/Users/hank/Downloads/GoDownloader/index.ts`, `/Users/hank/Downloads/GoDownloader/index.mp4`)
	if err != nil {
		t.Error(err)
	}
}
