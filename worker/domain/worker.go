package domain

type Worker struct {
	workingJobs []*Job
	capacity    int
}

func NewWorker(cap int) *Worker {
	return &Worker{capacity: cap}
}

func (w *Worker) AddJob(j *Job) {
	w.workingJobs = append(w.workingJobs, j)
}

func (w *Worker) ExecuteAllJob(interval int) int {
	point := 0
	newWorkingJob := []*Job{}
	for i := 0; i < len(w.workingJobs); i++ {
		j := w.workingJobs[i]

		currentPoint, done := j.Work(interval)
		point += currentPoint
		if !done {
			newWorkingJob = append(newWorkingJob, w.workingJobs[i])
		}
	}
	w.workingJobs = newWorkingJob
	return point
}

func (w *Worker) NumOfJob() int {
	return len(w.workingJobs)
}
