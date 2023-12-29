package api

import (
	"net/http"
	"path"

	"github.com/admpub/gohls-server/internal/config"
	"github.com/admpub/gohls-server/internal/hls"
)

type foldersHandler struct {
	conf *config.Config
}

func NewFoldersHandler(conf *config.Config) *foldersHandler {
	return &foldersHandler{conf}
}

func (s *foldersHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	videos := make([]*hls.ListResponseVideo, 0)
	folders := make([]*hls.ListResponseFolder, 0)
	parents := make([]*hls.ListResponseFolder, 0)
	response := &hls.ListResponse{Error: nil, Name: "Home", Path: "/", Parents: &parents, Folders: folders, Videos: videos}
	for _, f := range s.conf.Folders {
		folder := &hls.ListResponseFolder{Name: f.Title, Path: path.Join(r.URL.Path, f.Id)}
		folders = append(folders, folder)
	}
	response.Videos = videos
	response.Folders = folders
	hls.ServeJson(200, response, w)
}
