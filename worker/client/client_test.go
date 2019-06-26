package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/domain"
)

func TestClient_GetJob(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("time")

		if q == "00:00:01" {
			data, err := ioutil.ReadFile("../tests/sample_data/00001.job")
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "なし")
			}

			w.Write(data)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "なし")
		}
	}))

	defer testServer.Close()

	type fields struct {
		port int
		addr string
	}
	type args struct {
		t time.Time
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *domain.Job
		wantErr bool
	}{
		{
			name: "存在するJob",
			fields: fields{
				addr: testServer.URL,
			},
			args: args{
				t: time.Date(0, 1, 1, 0, 0, 1, 0, time.UTC),
			},
			want: &domain.Job{
				ID:       1,
				Created:  time.Date(0, 1, 1, 0, 00, 1, 0, time.UTC),
				Priority: domain.Low,
				Tasks:    []int{5, 6, 7},
			},
			wantErr: false,
		},
		{
			name: "存在しないJob",
			fields: fields{
				addr: testServer.URL,
			},
			args: args{
				t: time.Date(0, 1, 1, 5, 14, 10, 0, time.UTC),
			},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Client{
				addr: tt.fields.addr,
			}
			got, err := c.GetJob(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.GetJob() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.GetJob() = %v, want %v", got, tt.want)
			}
		})
	}
}
