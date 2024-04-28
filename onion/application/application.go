package application

import (
	"onion-architecture/domain"
	"onion-architecture/repository"
)

type JobApplication interface {
	FindAll() ([]*domain.Job, error)
	FindByID(id int) (*domain.Job, error)
	Create(job *domain.Job) error
	Update(job *domain.Job) error
	Delete(id int) error
}

type JobApplicationImpl struct {
	JobRepository repository.JobRepository
}

func NewJobApplication(jobRepository repository.JobRepository) JobApplication {
	return &JobApplicationImpl{
		JobRepository: jobRepository,
	}
}

func (a *JobApplicationImpl) FindAll() ([]*domain.Job, error) {
	return a.JobRepository.FindAll()
}

func (a *JobApplicationImpl) FindByID(id int) (*domain.Job, error) {
	return a.JobRepository.FindByID(id)
}

func (a *JobApplicationImpl) Create(job *domain.Job) error {
	if err := job.Validate(); err != nil {
		return err
	}
	return a.JobRepository.Save(job)
}

func (a *JobApplicationImpl) Update(job *domain.Job) error {
	if err := job.Validate(); err != nil {
		return err
	}
	return a.JobRepository.Update(job)
}

func (a *JobApplicationImpl) Delete(id int) error {
	return a.JobRepository.Delete(id)
}
