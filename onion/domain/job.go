package domain

import (
	"errors"
)

type Job struct {
	ID            int
	Name          string
	Content       string
	SecretContent string
}

func (j *Job) Validate() error {
	if j.Name == "" {
		return errors.New("name is required")
	}
	if j.Content == "" {
		return errors.New("content is required")
	}
	return nil
}

func NewJob(id int, name, content, secretContent string) (*Job, error) {
	j := &Job{
		ID:            id,
		Name:          name,
		Content:       content,
		SecretContent: secretContent,
	}
	if err := j.Validate(); err != nil {
		return nil, err
	}
	return j, nil
}
