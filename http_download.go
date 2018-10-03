package main

import (
	"net/http"
	"net/url"
	"path/filepath"
)

type downloadHandler struct {
	dir string
}

func NewDownloadHandler(root string) *downloadHandler {
	return &downloadHandler{root}
}

func (s *downloadHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	filename := url.QueryEscape(filepath.Base(r.URL.Path))
	w.Header().Set("Content-Disposition", "attachment; filename="+filename+"; filename*=UTF-8''"+filename)

	http.ServeFile(w, r, filepath.Join(s.dir, r.URL.Path))
}
