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

func (j *Job) Work(secs int) (point int, done bool) {
	for i := 0; i < secs; i++ {
		j.Tasks[j.CurrentTask]--

		if j.Tasks[j.CurrentTask] == 0 {
			j.CurrentTask++
		}

		//すべてのタスクが完了したので削除する (TODO 計算量がO(n)なので別の本心を考える)
		if j.CurrentTask == len(j.Tasks) {
			return 0, true
		}
	}
	return j.Tasks[j.CurrentTask], false

}
