package utils

import (
	"os"
	"strings"
	"testing"
)

func TestConvert(t *testing.T) {
	tsFile := `/Users/hank/Downloads/GoDownloader/index.ts`
	if _, err := os.Stat(tsFile); err != nil {
		return
	}
	err := ConvertToMP4(tsFile, strings.TrimSuffix(tsFile, `.ts`)+`.mp4`)
	if err != nil {
		t.Error(err)
	}
}
