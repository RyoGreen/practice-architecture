package postgresql

import (
	"architecture/onion/domain"
	"architecture/onion/repository"
	"database/sql"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Init() (err error) {
	db, err = sql.Open("postgres", "postgres://user:pass@host:port/dbname")
	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	db.Close()
}

type JobRepositoryPostgres struct {
}

func NewJobRepositoryPostgres() repository.JobRepository {
	return &JobRepositoryPostgres{}
}

func (r *JobRepositoryPostgres) FindAll() ([]*domain.Job, error) {
	rows, err := db.Query("SELECT * FROM jobs")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var jobs []*domain.Job
	for rows.Next() {
		var job domain.Job
		err := rows.Scan(&job.ID, &job.Name, &job.Content, &job.Salary)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, &job)
	}

	return jobs, nil
}

func (r *JobRepositoryPostgres) FindByID(id int) (*domain.Job, error) {
	row := db.QueryRow("SELECT * FROM jobs WHERE id = ?", id)
	var job domain.Job
	err := row.Scan(&job.ID, &job.Name, &job.Content, &job.Salary)
	if err != nil {
		return nil, err
	}
	return &job, nil
}

func (r *JobRepositoryPostgres) Save(job *domain.Job) error {
	if _, err := db.Exec("INSERT INTO jobs (name, content, salary) VALUES (?, ?, ?)", job.Name, job.Content, job.Salary); err != nil {
		return err
	}
	return nil
}

func (r *JobRepositoryPostgres) Update(job *domain.Job) error {
	if _, err := db.Exec("UPDATE jobs SET name = ?, content = ?, salary = ? WHERE id = ?", job.Name, job.Content, job.Salary, job.ID); err != nil {
		return err
	}
	return nil
}

func (r *JobRepositoryPostgres) Delete(id int) error {
	if _, err := db.Exec("DELETE FROM jobs WHERE id = ?", id); err != nil {
		return err
	}
	return nil
}
