package domain

type JobNotFoundError struct {
}

func (e *JobNotFoundError) Error() string {
	return "job not found"
}
