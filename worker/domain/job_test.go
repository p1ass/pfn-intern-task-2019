package domain

import (
	"testing"
	"time"
)

func TestJob_Work(t *testing.T) {
	type fields struct {
		ID          int
		Created     time.Time
		Priority    Priority
		Tasks       []int
		CurrentTask int
	}
	tests := []struct {
		name      string
		fields    fields
		wantPoint int
		wantDone  bool
	}{
		{
			name: "1つのタスクが丁度終わる",
			fields: fields{
				ID:          0,
				Created:     time.Time{},
				Priority:    Low,
				Tasks:       []int{1},
				CurrentTask: 0,
			},
			wantPoint: 0,
			wantDone:  true,
		}, {
			name: "1秒進める",
			fields: fields{
				ID:          0,
				Created:     time.Time{},
				Priority:    Low,
				Tasks:       []int{5, 4, 3},
				CurrentTask: 0,
			},
			wantPoint: 4,
			wantDone:  false,
		},
		{
			name: "1つめのタスクが丁度終わる",
			fields: fields{
				ID:          0,
				Created:     time.Time{},
				Priority:    Low,
				Tasks:       []int{1, 4, 3},
				CurrentTask: 0,
			},
			wantPoint: 4,
			wantDone:  false,
		},
		{
			name: "すべてのタスクが丁度完了する",
			fields: fields{
				ID:          0,
				Created:     time.Time{},
				Priority:    Low,
				Tasks:       []int{1},
				CurrentTask: 0,
			},
			wantPoint: 0,
			wantDone:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			j := &Job{
				ID:          tt.fields.ID,
				Created:     tt.fields.Created,
				Priority:    tt.fields.Priority,
				Tasks:       tt.fields.Tasks,
				CurrentTask: tt.fields.CurrentTask,
			}
			gotPoint, gotDone := j.Work()
			if gotPoint != tt.wantPoint {
				t.Errorf("Job.Work() gotPoint = %v, wantPoint %v", gotPoint, tt.wantPoint)
			}
			if gotDone != tt.wantDone {
				t.Errorf("Job.Work() gotDone = %v, wantPoint %v", gotDone, tt.wantDone)
			}
		})
	}
}
