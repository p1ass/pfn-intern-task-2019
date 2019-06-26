package domain

type Worker struct {
	workingJobs  []*Job
	jobQueue     []*Job
	currentPoint int
	capacity     int
}

func NewWorker(cap int) *Worker {
	return &Worker{capacity: cap}
}

func (w *Worker) AddJob(j *Job) {
	if w.currentPoint+j.Tasks[j.CurrentTask] > w.capacity {
		w.addJobToQueue(j)
	} else {
		w.workingJobs = append(w.workingJobs, j)
	}
}

func (w *Worker) addJobToQueue(j *Job) {
	w.jobQueue = append(w.jobQueue, j)
}

func (w *Worker) ExecuteAllJob(secs int) int {
	sumPoint := 0

	for i := 0; i < secs; i++ {
		sumPoint = 0

		newWorkingJob := []*Job{}

		for i := 0; i < len(w.workingJobs); i++ {
			j := w.workingJobs[i]

			point, done := j.Work()

			if sumPoint+point > w.capacity {
				w.addJobToQueue(j)
			} else {
				sumPoint += point
				if !done {
					newWorkingJob = append(newWorkingJob, j)
				}
			}
		}

		w.workingJobs = newWorkingJob
		w.currentPoint = sumPoint
		w.currentPoint += w.moveJobToWorking()
	}
	return w.currentPoint
}

func (w *Worker) moveJobToWorking() int {
	point := 0
	for len(w.jobQueue) > 0 {
		j := w.jobQueue[0]
		if w.currentPoint+j.Tasks[j.CurrentTask] <= w.capacity {
			w.workingJobs = append(w.workingJobs, j)
			w.jobQueue = w.jobQueue[1:]
			point += j.Tasks[j.CurrentTask]
		} else {
			break
		}
	}
	return point
}

func (w *Worker) NumOfJob() int {
	return len(w.workingJobs)
}
