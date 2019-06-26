package server

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestServer_GetJobHandler(t *testing.T) {
	type fields struct {
		time string
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "Jobが存在する時間",
			fields: fields{
				time: "00:00:01",
			},
			want: getWantBytes("../tests/sample_data/00001.job", t),
		}, {
			name: "Jobが存在しない時間",
			fields: fields{
				time: "00:00:02",
			},
			want: []byte("なし"),
		},
	}

	s := NewServer("../tests/sample_data", 0)
	testServer := httptest.NewServer(http.HandlerFunc(s.GetJobHandler))
	defer testServer.Close()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := url.Values{}
			q.Add("time", tt.fields.time)
			resp, err := http.Get(testServer.URL + "?" + q.Encode())
			if err != nil {
				t.Fatalf("failed to request to server: %s", err)
			}
			defer resp.Body.Close()
			b, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				t.Fatalf("Failed to read response body: %s", err)
			}
			if !reflect.DeepEqual(b, tt.want) {
				t.Errorf("Test_GetJob() response body = %s, want %s", b, tt.want)
			}
		})
	}
}

func getWantBytes(filename string, t *testing.T) []byte {
	t.Helper()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatalf("failed to read wanted file: %s", err)
	}
	return data
}
