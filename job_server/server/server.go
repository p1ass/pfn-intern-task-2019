package server

import (
	"bufio"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

type Server struct {
	dir      string
	timeMemo map[string]string
	port     int
}

func NewServer(dir string, port int) *Server {
	s := &Server{dir, map[string]string{}, port}
	s.readJob()
	return s
}

func (s *Server) Start() error {
	http.HandleFunc("/", s.GetJobHandler)
	log.Println("Starting server...")
	return http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
}

// key が Created、 value が ファイル名のハッシュマップを作成し、リクエストごとのファイルの探索をO(1)で出来るようにする
func (s *Server) readJob() {
	files, err := ioutil.ReadDir(s.dir)
	if err != nil {
		log.Fatalf("failed to read directory : %s", err) // jobを読み込めなかったらサーバを落としてよい。
	}

	for _, f := range files {
		fp, err := os.Open(filepath.Join(s.dir, f.Name()))
		if err != nil {
			log.Fatalf("failed to read job files : %s", err) // jobを読み込めなかったらサーバを落としてよい。
		}
		defer fp.Close()

		scanner := bufio.NewScanner(fp)

		for scanner.Scan() {
			if scanner.Text() == "[Created]" {
				scanner.Scan()
				s.timeMemo[scanner.Text()] = f.Name()
			}
		}

		if err = scanner.Err(); err != nil {
			log.Fatalf("failed to read job files : %s", err) // jobを読み込めなかったらサーバを落としてよい。
		}
	}

}
