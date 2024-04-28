package domain

import "errors"

type Job struct {
	ID      int
	Name    string
	Content string
	Salary  int
}

func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("name is required")
	}
	if j.Content == "" {
		return errors.New("content is required")
	}
	if j.Salary > 0 {
		return errors.New("salary must be greater than 0")
	}
	return nil
}
