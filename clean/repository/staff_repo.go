package repository

import (
	"architecture/clean/entity"
	"errors"
)

type StaffRepository interface {
	List() ([]*entity.Staff, error)
	Get(id int) (*entity.Staff, error)
	Create(*entity.Staff) (*entity.Staff, error)
	Update(*entity.Staff) (*entity.Staff, error)
	Delete(id int) error
}

type staffRepositoryMap struct {
}

func NewStaffRepositoryMap() StaffRepository {
	return &staffRepositoryMap{}
}

var staffsMap = map[int]entity.Staff{}

var (
	ErrStaffNotFound = errors.New("Staff not found")
	ErrStaffExists   = errors.New("Staff already exists")
)

func (r *staffRepositoryMap) List() ([]*entity.Staff, error) {
	staffs := make([]*entity.Staff, 0, len(staffsMap))
	for _, v := range staffs {
		staffs = append(staffs, v)
	}
	return staffs, nil
}

func (r *staffRepositoryMap) Get(id int) (*entity.Staff, error) {
	if v, ok := staffsMap[id]; ok {
		return &v, nil
	}
	return nil, ErrStaffNotFound
}

func (r *staffRepositoryMap) Create(staff *entity.Staff) (*entity.Staff, error) {
	if _, ok := staffsMap[staff.ID]; ok {
		return nil, ErrStaffExists
	}
	staffsMap[staff.ID] = *staff
	return staff, nil
}

func (r *staffRepositoryMap) Update(staff *entity.Staff) (*entity.Staff, error) {
	if _, ok := staffsMap[staff.ID]; !ok {
		return nil, ErrStaffNotFound
	}
	staffsMap[staff.ID] = *staff
	return staff, nil
}

func (r *staffRepositoryMap) Delete(id int) error {
	if _, ok := staffsMap[id]; !ok {
		return ErrStaffNotFound
	}
	delete(staffsMap, id)
	return nil
}
