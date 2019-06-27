package domain

import (
	"container/heap"
)

type Worker struct {
	workingJobs  []*Job
	jobPQ        *JobPriorityQueue
	currentPoint int
	capacity     int
}

func NewWorker(cap int) *Worker {
	pq := NewJobPriorityQueue()
	heap.Init(pq)

	return &Worker{jobPQ: pq, capacity: cap}
}

//AddJob はジョブを実行中かキューに追加する関数です。
func (w *Worker) AddJob(j *Job) {
	if w.currentPoint+j.Tasks[j.CurrentTask] > w.capacity || w.jobPQ.Len() != 0 {
		w.addJobToQueue(j)
	} else {
		w.addJobToWorking(j)
	}
}

func (w *Worker) addJobToWorking(j *Job) {
	w.workingJobs = append(w.workingJobs, j)
	w.currentPoint += j.Tasks[j.CurrentTask]
}

func (w *Worker) addJobToQueue(j *Job) {
	heap.Push(w.jobPQ, j)
}

//ExecuteAllJob は実行できるすべてのジョブを与えられた秒数だけ実行します。
func (w *Worker) ExecuteAllJob(secs int) int {
	sumPoint := 0

	for i := 0; i < secs; i++ {
		sumPoint = 0

		newWorkingJob := []*Job{}

		for i := 0; i < len(w.workingJobs); i++ {
			j := w.workingJobs[i]

			point, done := j.Work()

			if !j.AlreadyCompleted() {
				if done {
					w.addJobToQueue(j)
				} else {
					sumPoint += point
					newWorkingJob = append(newWorkingJob, j)
				}
			}
		}

		w.workingJobs = newWorkingJob
		w.currentPoint = sumPoint

		w.updatePQ()
		w.moveJobToWorking()
	}
	return w.currentPoint
}

func (w *Worker) CurrentPoint() int {
	return w.currentPoint
}

// moveJobToWorking は空いているキャパシティに応じて、新しいジョブを実行中に移す関数です。
func (w *Worker) moveJobToWorking() int {
	point := 0
	for w.jobPQ.Len() > 0 {
		if j, ok := heap.Pop(w.jobPQ).(*Job); ok {
			if w.currentPoint+j.Tasks[j.CurrentTask] <= w.capacity {
				w.addJobToWorking(j)
				point += j.Tasks[j.CurrentTask]
			} else {
				heap.Push(w.jobPQ, j)
				break
			}
		}
	}
	return point
}

func (w *Worker) NumOfJob() int {
	return len(w.workingJobs)
}

func (w *Worker) updatePQ() {
	w.jobPQ.SetUnusedCap(w.capacity - w.currentPoint)
	heap.Init(w.jobPQ)
}
