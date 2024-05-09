package application

import (
	"crypto/sha256"
	"fmt"
	"onion-architecture/domain"
	"onion-architecture/presentation/request"
	"onion-architecture/presentation/response"
	"onion-architecture/repository"
)

type JobApplicationImpl struct {
	JobRepository repository.JobRepository
}

func NewJobApplication(jobRepository repository.JobRepository) *JobApplicationImpl {
	return &JobApplicationImpl{
		JobRepository: jobRepository,
	}
}

func (a *JobApplicationImpl) FindAll() ([]*response.JobResponse, error) {
	jobs, err := a.JobRepository.FindAll()
	if err != nil {
		return nil, err
	}
	var jobResponse []*response.JobResponse

	for _, job := range jobs {
		jobResponse = append(jobResponse, &response.JobResponse{
			ID:      job.ID,
			Name:    job.Name,
			Content: job.Content,
		})
	}
	return jobResponse, nil
}

type JobDetail struct {
	ID            int
	Name          string
	Content       string
	SecretContent string
}

func (d *JobDetail) Hash() {
	key := fmt.Sprintf("%v%s%s%s", d.ID, d.Name, d.Content, d.SecretContent)
	h := sha256.Sum256([]byte(key))
	d.SecretContent = string(h[:])
}

func (a *JobApplicationImpl) FindByID(id int) (*response.JobResponseWithSecret, error) {
	job, err := a.JobRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	jobDetail := &JobDetail{
		ID:            job.ID,
		Name:          job.Name,
		Content:       job.Content,
		SecretContent: job.SecretContent,
	}
	jobDetail.Hash()

	return &response.JobResponseWithSecret{
		ID:            jobDetail.ID,
		Content:       jobDetail.Content,
		Name:          jobDetail.Name,
		SecretContent: jobDetail.SecretContent,
	}, nil
}

func (a *JobApplicationImpl) Create(job *request.JobRequest) error {
	j, err := domain.NewJob(job.ID, job.Name, job.Content, job.SecretContent)
	if err != nil {
		return err
	}

	return a.JobRepository.Save(j)
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
