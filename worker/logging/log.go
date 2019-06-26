package logging

import (
	"fmt"
	"time"
)

type Log struct {
	timestamp time.Time
	Point     int
}

type Logger struct {
	data []*Log
}

func NewLogger() *Logger {
	return &Logger{}
}
func (l *Logger) Add(log *Log) {
	l.data = append(l.data, log)
}

func NewLog(timestamp time.Time, point int) *Log {
	return &Log{timestamp, point}
}

func (l *Logger) Print() {
	for _, l := range l.data {
		fmt.Printf("%s, %d\n", l.timestamp.Format("15:04:05"), l.Point)
	}
}
