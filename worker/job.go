package main

import "time"

type Priority int

const (
	Low Priority = iota
	High
)

type Job struct {
	ID       int
	Created  time.Time
	Priority Priority
	Tasks    []int
}
