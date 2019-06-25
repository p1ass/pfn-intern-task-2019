package main

import (
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

	logs := []*Log{}

	cli := client.NewClient("http://localhost:8080")

	current := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	interval := 1

	worker := domain.NewWorker()

outer:
	for {
		point := worker.ExecuteAllJob(interval)

		newJob, err := cli.GetJob(current)
		if err != nil {
			switch e := err.(type) {
			case *domain.JobNotFoundError:
				//TODO 送られてきたデータを見て時間のインクリメントをやめる処理を実装する
				if current == time.Date(0, 1, 1, 1, 59, 59, 0, time.UTC) {
					break outer
				}
			default:
				log.Fatalf("failed to get job from server: %s", e) //予期しないエラーなのでワーカーを落とす
			}
		}

		if newJob != nil {
			worker.AddJob(newJob)
			point += newJob.Tasks[0]
		}

		l := &Log{current, point}
		logs = append(logs, l)
		current = current.Add(time.Duration(interval) * time.Second)
	}

	for _, l := range logs {
		fmt.Printf("%s, %d\n", l.timestamp.Format("15:04:05"), l.Point)
	}
}
