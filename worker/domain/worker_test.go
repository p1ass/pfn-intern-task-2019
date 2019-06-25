package domain

import (
	"testing"
	"time"
)

func TestWorker_ExecuteAllJob(t *testing.T) {
	type fields struct {
		workingJobs []*Job
	}
	type args struct {
		interval int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantPoint int
		wantNum   int
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
			},
			args: args{
				interval: 1,
			},
			wantPoint: 3,
			wantNum:   2,
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
			},
			args: args{
				interval: 1,
			},
			wantPoint: 2,
			wantNum:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Worker{
				workingJobs: tt.fields.workingJobs,
			}
			if got := w.ExecuteAllJob(tt.args.interval); got != tt.wantPoint {
				t.Errorf("Worker.ExecuteAllJob() = %v, wantPoint=%v", got, tt.wantPoint)
			}

			if len(w.workingJobs) != tt.wantNum {
				t.Errorf("len(workingJobs) = %v, wantNum=%v", len(w.workingJobs), tt.wantNum)
			}
		})
	}
}
