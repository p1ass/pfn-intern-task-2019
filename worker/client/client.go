package client

import (
	"bufio"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/naoki-kishi/pfn-intern-task-2019/worker/domain"
)

type Client struct {
	addr string
}

func NewClient(addr string) *Client {
	return &Client{addr}
}

func (c *Client) GetJob(t time.Time) (*domain.Job, error) {

	timeStr := t.Format("15:04:05")
	query := url.Values{}
	query.Add("time", timeStr)
	resp, err := http.Get(c.addr + "?" + query.Encode())
	if err != nil {
		return nil, fmt.Errorf("failed to request job: %s", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return nil, &domain.JobNotFoundError{}
	} else if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("response is not OK status=%d", resp.StatusCode)
	}

	defer resp.Body.Close()
	scanner := bufio.NewScanner(resp.Body)

	job := &domain.Job{}

	for scanner.Scan() {
		switch scanner.Text() {
		case "[JobID]":
			scanner.Scan()
			idStr := scanner.Text()
			id, err := strconv.Atoi(idStr)
			if err != nil {
				return nil, fmt.Errorf("failed to paser job id: %s", err)
			}
			job.ID = id

		case "[Created]":
			scanner.Scan()
			timeStr := scanner.Text()
			t, err := time.Parse("15:04:05", timeStr)
			if err != nil {
				return nil, fmt.Errorf("failed to paser created time: %s", err)
			}
			job.Created = t

		case "[Priority]":
			scanner.Scan()
			p := scanner.Text()
			if p == "Low" {
				job.Priority = domain.Low
			} else if p == "High" {
				job.Priority = domain.High
			}

		case "[Tasks]":
			for scanner.Scan() {
				taskStr := scanner.Text()
				task, err := strconv.Atoi(taskStr)
				if err != nil {
					return nil, fmt.Errorf("failed to paser task: %s", err)
				}
				job.Tasks = append(job.Tasks, task)
			}
		}

	}
	return job, nil
}
