package domain

type JobPriorityQueue []*Job

func NewJobPriorityQueue() *JobPriorityQueue {
	return &JobPriorityQueue{}
}

func (pq JobPriorityQueue) Less(i, j int) bool {
	if pq[i].Priority > pq[j].Priority {
		return true
	} else if pq[i].Priority < pq[j].Priority {
		return false
	} else {

		if pq[i].Created.Before(pq[j].Created) {
			return true
		}
		return false
	}
}

func (pq JobPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq JobPriorityQueue) Len() int {
	return len(pq)
}

func (pq *JobPriorityQueue) Push(val interface{}) {
	if v, ok := val.(*Job); ok {
		*pq = append(*pq, v)
	}
}

func (pq *JobPriorityQueue) Pop() interface{} {
	old := *pq
	size := len(*pq)
	job := old[size-1]
	*pq = old[:size-1]
	return job
}
