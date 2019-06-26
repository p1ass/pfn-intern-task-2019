package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/domain"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/client"
)

type Log struct {
	timestamp time.Time
	Point     int
}

func main() {
	capacity := flag.Int("c", 15, "Capacity")
	port := flag.String("p", "8080", "Server port number")
	flag.Parse()

	logs := []*Log{}
	cli := client.NewClient("http://localhost:" + *port)

	current := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	interval := 1

	worker := domain.NewWorker(*capacity)

outer:
	for {
		worker.ExecuteAllJob(interval)

		newJob, err := cli.GetJob(current)
		if err != nil {
			switch e := err.(type) {
			case *domain.JobNotFoundError:
				//TODO 送られてきたデータを見て時間のインクリメントをやめる処理を実装する
				if current == time.Date(0, 1, 1, 1, 29, 59, 0, time.UTC) {
					break outer
				}
			default:
				log.Fatalf("failed to get job from server: %s", e) //予期しないエラーなのでワーカーを落とす
			}
		}

		if newJob != nil {
			worker.AddJob(newJob)
		}

		l := &Log{current, worker.CurrentPoint()}
		logs = append(logs, l)
		current = current.Add(time.Duration(interval) * time.Second)
	}

	for _, l := range logs {
		fmt.Printf("%s, %d\n", l.timestamp.Format("15:04:05"), l.Point)
	}
}
