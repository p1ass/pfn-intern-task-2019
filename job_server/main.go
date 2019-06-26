package main

import (
	"flag"
	"log"

	"github.com/naoki-kishi/pfn-intern-task-2019/job_server/server"
)

func main() {
	port := flag.Int("p", 8080, "Server port number")
	dir := flag.String("d", "./data", "Job data directory")
	flag.Parse()

	s := server.NewServer(*dir, *port)

	log.Fatal(s.Start())
}
