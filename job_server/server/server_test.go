package server

import (
	"context"
	"os/exec"
	"testing"
)

func TestServer_readJob(t *testing.T) {
	type fields struct {
		dir      string
		timeMemo map[string]string
		port     int
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]string
	}{
		{
			name: "ジョブデータを正しく読み込める",
			fields: fields{
				dir:      "../tests/sample_data",
				timeMemo: map[string]string{},
				port:     0,
			},
			want: map[string]string{"05:13:10": "test_00001.job", "00:00:01s": "test_00003.job"},
		},
	}
	ctx := context.Background()
	cmd := exec.CommandContext(ctx, "make", "test-server")

	if err := cmd.Start(); err != nil {
		t.Fatalf("failed to start test server: %s", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Server{
				dir:      tt.fields.dir,
				timeMemo: tt.fields.timeMemo,
				port:     tt.fields.port,
			}
			s.readJob()
		})
	}
}
