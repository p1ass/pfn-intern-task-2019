package main

import (
	"log"

	"github.com/naoki-kishi/pfn-intern-task-2019/job_server/server"
)

func main() {
	s := server.NewServer("./data", 8080)

	log.Fatal(s.Start())
}
