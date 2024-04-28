package mysql

import (
	"architecture/clean/entity"
	"architecture/clean/repository"

	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return err
	}
	c := mysql.Config{
		DBName:    "clean_architecture",
		User:      "user",
		Passwd:    "password",
		Addr:      "localhost:3306",
		Net:       "tcp",
		ParseTime: true,
		Loc:       jst,
	}
	db, err = sql.Open("mysql", c.FormatDSN())
	if err != nil {
		return err
	}
	return nil
}

type staffRepositoryMysql struct {
}

func NewStaffRepositoryMysql(config *mysql.Config) repository.StaffRepository {
	return &staffRepositoryMysql{}
}

func (r *staffRepositoryMysql) List() ([]*entity.Staff, error) {
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

func (r *staffRepositoryMysql) Get(id int) (*entity.Staff, error) {
	query := "SELECT id, name, email FROM staff WHERE id = ?"
	row := db.QueryRow(query, id)
	var staff entity.Staff
	if err := row.Scan(&staff.ID, &staff.Name, &staff.Email); err != nil {
		return nil, err
	}
	return &staff, nil
}

func (r *staffRepositoryMysql) Create(staff *entity.Staff) (*entity.Staff, error) {
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

func (r *staffRepositoryMysql) Update(staff *entity.Staff) (*entity.Staff, error) {
	query := "UPDATE staff SET name = ?, email = ? WHERE id = ?"
	_, err := db.Exec(query, staff.Name, staff.Email, staff.ID)
	if err != nil {
		return nil, err
	}
	return staff, nil
}

func (r *staffRepositoryMysql) Delete(id int) error {
	query := "DELETE FROM staff WHERE id = ?"
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}
