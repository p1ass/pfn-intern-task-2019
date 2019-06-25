package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestClient_GetJob(t *testing.T) {
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query().Get("time")

		if q == "05:13:10" {
			data, err := ioutil.ReadFile("./tests/sample_data/test_00001.job")
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
		want    *Job
		wantErr bool
	}{
		{
			name: "存在するJob",
			fields: fields{
				addr: testServer.URL,
			},
			args: args{
				t: time.Date(0, 1, 1, 5, 13, 10, 0, time.UTC),
			},
			want: &Job{
				ID:       0,
				Created:  time.Date(0, 1, 1, 5, 13, 10, 0, time.UTC),
				Priority: Low,
				Tasks:    []int{3, 8, 10, 1},
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
