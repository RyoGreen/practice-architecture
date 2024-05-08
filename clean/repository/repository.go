package repository

import (
	"clean-architecture/entity"
	"errors"
)

type StaffRepository interface {
	List() ([]*entity.Staff, error)
	Get(id int) (*entity.Staff, error)
	Create(*entity.Staff) (*entity.Staff, error)
	Update(*entity.Staff) (*entity.Staff, error)
	Delete(id int) error
}

var (
	ErrStaffNotFound = errors.New("staff not found")
	ErrStaffExists   = errors.New("staff already exists")
)
