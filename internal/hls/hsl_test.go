package hls

import "testing"

func TestConvertMP4(t *testing.T) {
	return
	err := ConvertToMP4(`/Users/hank/Downloads/hifilm.mov`, `/Users/hank/Downloads/hifilm.mp4`)
	if err != nil {
		t.Error(err)
	}
}
