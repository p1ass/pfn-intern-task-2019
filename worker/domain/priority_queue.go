package domain

type JobPriorityQueue struct {
	data      []*Job
	unusedCap int
}

func NewJobPriorityQueue() *JobPriorityQueue {
	return &JobPriorityQueue{}
}

func (pq JobPriorityQueue) Less(i, j int) bool {
	job1 := pq.data[i]
	job2 := pq.data[j]

	//優先順位が高いものを先に処理する
	if job1.Priority > job2.Priority {
		return true
	} else if job1.Priority < job2.Priority {
		return false
	}

	// 終戦順位が同じもの同士のときは、空いているキャパシティをなるべく埋めるようにする
	task1 := job1.Tasks[job1.CurrentTask]
	task2 := job2.Tasks[job2.CurrentTask]
	if task1 <= pq.unusedCap && task2 > pq.unusedCap {
		return true
	} else if task1 > pq.unusedCap && task2 <= pq.unusedCap {
		return false
	} else {
		if task1 < task2 {
			return true
		} else {
			return false
		}
	}
}

func (pq JobPriorityQueue) Swap(i, j int) {
	pq.data[i], pq.data[j] = pq.data[j], pq.data[i]
}

func (pq JobPriorityQueue) Len() int {
	return len(pq.data)
}

func (pq *JobPriorityQueue) Push(val interface{}) {
	if v, ok := val.(*Job); ok {
		pq.data = append(pq.data, v)
	}
}

func (pq *JobPriorityQueue) Pop() interface{} {
	size := len(pq.data)
	job := pq.data[size-1]
	pq.data = pq.data[:size-1]
	return job
}

func (pq *JobPriorityQueue) SetUnusedCap(unusedCap int) {
	pq.unusedCap = unusedCap
}
