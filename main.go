package main

import (
	"log"

	"github.com/naoki-kishi/pfn-intern-task-2019/server"
)

func main() {
	s := server.NewServer("./data", 8080)

	log.Fatal(s.Start())
}
