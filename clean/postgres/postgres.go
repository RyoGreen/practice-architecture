package postgres

import (
	"clean-architecture/entity"
	"clean-architecture/repository"
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

type staffRepositoryPostgres struct {
}

func NewStaffRepositoryPostgres() repository.StaffRepository {
	return &staffRepositoryPostgres{}
}

func (r *staffRepositoryPostgres) List() ([]*entity.Staff, error) {
	query := "SELECT id, name, email FROM staff"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var staffs []*entity.Staff
	for rows.Next() {
		var staff entity.Staff
		if err := rows.Scan(&staff.ID, &staff.Name, &staff.Email); err != nil {
			return nil, err
		}
		staffs = append(staffs, &staff)
	}
	return staffs, nil
}

func (r *staffRepositoryPostgres) Get(id int) (*entity.Staff, error) {
	query := "SELECT id, name, email FROM staff WHERE id = ?"
	row := db.QueryRow(query, id)
	var staff entity.Staff
	if err := row.Scan(&staff.ID, &staff.Name, &staff.Email); err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepositoryPostgres) Create(staff *entity.Staff) (*entity.Staff, error) {
	query := "INSERT INTO staff (name, email) VALUES (?, ?)"
	result, err := db.Exec(query, staff.Name, staff.Email)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	staff.ID = int(id)
	return staff, nil
}

func (r *staffRepositoryPostgres) Update(staff *entity.Staff) (*entity.Staff, error) {
	query := "UPDATE staff SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, staff.Name, staff.Email, staff.ID)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepositoryPostgres) Delete(id int) error {
	query := "DELETE FROM staff WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
