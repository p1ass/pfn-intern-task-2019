package server

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
)

func (s *Server) GetJobHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("time")

		if fileName, ok := s.timeMemo[q]; ok {
			data, err := ioutil.ReadFile(filepath.Join(s.dir, fileName))
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "なし")
			}

			w.Write(data)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "なし")
		}
	}
}
