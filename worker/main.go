package main

import (
	"flag"
	"log"
	"time"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/logging"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/domain"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/client"
)

func main() {
	capacity := flag.Int("c", 15, "Capacity")
	port := flag.String("p", "8080", "Server port number")
	flag.Parse()

	logger := logging.NewLogger()
	cli := client.NewClient("http://localhost:" + *port)
	worker := domain.NewWorker(*capacity)

	current := time.Date(0, 1, 1, 0, 0, 0, 0, time.UTC)
	interval := 1

	//時間を進めるループ
outer:
	for {
		worker.ExecuteAllJob(interval)

		newJob, err := cli.GetJob(current)
		if err != nil {
			switch e := err.(type) {
			case *domain.JobNotFoundError:
				//ワーカーを止める条件に関する記載がなかったので、とりあえず1時半で終了しています。
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

		logger.Add(logging.NewLog(current, worker.CurrentPoint()))
		current = current.Add(time.Duration(interval) * time.Second)
	}

	logger.Print()
}
