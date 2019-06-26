package domain

import (
	"container/heap"
	"testing"
	"time"
)

func TestWorker_ExecuteAllJob(t *testing.T) {
	type fields struct {
		workingJobs  []*Job
		jobQueue     []*Job
		jobPQ        *JobPriorityQueue
		currentPoint int
		capacity     int
	}
	type args struct {
		secs int
	}
	tests := []struct {
		name         string
		fields       fields
		args         args
		wantPoint    int
		wantNum      int
		wantQueueNum int
	}{
		{
			name: "1秒経過したとき正しいポイントが返ってくる",
			fields: fields{
				workingJobs: []*Job{
					{
						ID:          0,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{2},
						CurrentTask: 0,
					},
					{
						ID:          1,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{3, 1},
						CurrentTask: 0,
					},
				},
				jobQueue:     []*Job{},
				currentPoint: 5,
				capacity:     10000000,
			},
			args: args{
				secs: 1,
			},
			wantPoint:    3,
			wantNum:      2,
			wantQueueNum: 0,
		},
		{
			name: "2秒経過したとき正しいポイントが返ってくる",
			fields: fields{
				workingJobs: []*Job{
					{
						ID:          0,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{3},
						CurrentTask: 0,
					},
					{
						ID:          1,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{3, 1},
						CurrentTask: 0,
					},
				},
				jobQueue:     []*Job{},
				currentPoint: 6,
				capacity:     10000000,
			},
			args: args{
				secs: 2,
			},
			wantPoint:    2,
			wantNum:      2,
			wantQueueNum: 0,
		},
		{
			name: "完了したJobはwokingJobから消える",
			fields: fields{
				workingJobs: []*Job{
					{
						ID:          0,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{1},
						CurrentTask: 0,
					},
					{
						ID:          1,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{3, 1},
						CurrentTask: 0,
					},
				},
				jobQueue:     []*Job{},
				currentPoint: 4,
				capacity:     10000000,
			},
			args: args{
				secs: 1,
			},
			wantPoint:    2,
			wantNum:      1,
			wantQueueNum: 0,
		},
		{
			name: "capacityを超えたときjobQueueにjobが移動する",
			fields: fields{
				workingJobs: []*Job{
					{
						ID:          0,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{1, 10},
						CurrentTask: 0,
					},
					{
						ID:          1,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{3, 1},
						CurrentTask: 0,
					},
				},
				jobQueue:     []*Job{},
				currentPoint: 4,
				capacity:     10,
			},
			args: args{
				secs: 1,
			},
			wantPoint:    10,
			wantNum:      1,
			wantQueueNum: 1,
		},
		{
			name: "capacityの枠が空いた時jobQueueのjobがworkingJobに移動する",
			fields: fields{
				workingJobs: []*Job{
					{
						ID:          0,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{8},
						CurrentTask: 0,
					},
				},
				jobQueue: []*Job{
					{
						ID:          1,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{3, 1},
						CurrentTask: 0,
					},
				},
				currentPoint: 8,
				capacity:     10,
			},
			args: args{
				secs: 1,
			},
			wantPoint:    10,
			wantNum:      2,
			wantQueueNum: 0,
		},

		{
			name: "優先度が高いjobが先にworkingJobに移動する",
			fields: fields{
				workingJobs: []*Job{
					{
						ID:          0,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{8},
						CurrentTask: 0,
					},
				},
				jobQueue: []*Job{
					{
						ID:          1,
						Created:     time.Time{},
						Priority:    0,
						Tasks:       []int{4, 1},
						CurrentTask: 0,
					}, {
						ID:          1,
						Created:     time.Time{},
						Priority:    1,
						Tasks:       []int{3, 1},
						CurrentTask: 0,
					},
				},
				currentPoint: 8,
				capacity:     10,
			},
			args: args{
				secs: 1,
			},
			wantPoint:    10,
			wantNum:      2,
			wantQueueNum: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pq := NewJobPriorityQueue()
			heap.Init(pq)
			for _, j := range tt.fields.jobQueue {
				heap.Push(pq, j)
			}
			w := &Worker{
				workingJobs:  tt.fields.workingJobs,
				jobPQ:        pq,
				currentPoint: tt.fields.currentPoint,
				capacity:     tt.fields.capacity,
			}
			if got := w.ExecuteAllJob(tt.args.secs); got != tt.wantPoint {
				t.Errorf("Worker.ExecuteAllJob() = %v, wantPoint=%v", got, tt.wantPoint)
			}

			if len(w.workingJobs) != tt.wantNum {
				t.Errorf("len(workingJobs) = %v, wantNum=%v", len(w.workingJobs), tt.wantNum)
			}

			if w.jobPQ.Len() != tt.wantQueueNum {
				t.Errorf("w.jobPQ.Len() = %v, wantQueueNum=%v", w.jobPQ.Len(), tt.wantQueueNum)
			}
		})
	}
}
