package domain

import (
	"time"
)

type Priority int

const (
	Low Priority = iota
	High
)

type Job struct {
	ID          int
	Created     time.Time
	Priority    Priority
	Tasks       []int
	CurrentTask int
}

func (j *Job) Work() (point int, done bool) {
	j.Tasks[j.CurrentTask]--

	if j.Tasks[j.CurrentTask] == 0 {
		j.CurrentTask++
		return 0, true
	}

	return j.Tasks[j.CurrentTask], false
}

func (j *Job) AlreadyCompleted() bool {
	if j.CurrentTask == len(j.Tasks) {
		return true
	}
	return false
}
