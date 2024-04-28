package repository

import "onion-architecture/domain"

type JobRepository interface {
	FindAll() ([]*domain.Job, error)
	FindByID(id int) (*domain.Job, error)
	Save(job *domain.Job) error
	Update(job *domain.Job) error
	Delete(id int) error
}
