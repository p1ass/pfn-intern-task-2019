package domain

import (
	"time"
)

type Priority int

// 優先度を Enum のように扱えるように定数を宣言
const (
	Low Priority = iota
	High
)

//Job はジョブを表す構造体です。
type Job struct {
	ID          int
	Created     time.Time
	Priority    Priority
	Tasks       []int
	CurrentTask int
}

//Work はジョブを実行し、現在のタスクのポイントを返します。丁度タスクが完了したときは done を trueにします。
func (j *Job) Work() (point int, done bool) {
	j.Tasks[j.CurrentTask]--

	if j.Tasks[j.CurrentTask] == 0 {
		j.CurrentTask++
		return 0, true
	}

	return j.Tasks[j.CurrentTask], false
}

//AlreadyCompleted はすべてのタスクが完了しているかを bool で返す関数です。
func (j *Job) AlreadyCompleted() bool {
	if j.CurrentTask == len(j.Tasks) {
		return true
	}
	return false
}
