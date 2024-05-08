package dbmap

import (
	"clean-architecture/entity"
	"clean-architecture/repository"
)

type staffRepositoryMap struct {
}

func NewStaffRepositoryMap() entity.StaffRepository {
	return &staffRepositoryMap{}
}

var staffsMap = map[int]*entity.Staff{}

func (r *staffRepositoryMap) List() ([]*entity.Staff, error) {
	staffs := make([]*entity.Staff, 0, len(staffsMap))
	for _, v := range staffsMap {
		staffs = append(staffs, v)
	}
	return staffs, nil
}

func (r *staffRepositoryMap) Get(id int) (*entity.Staff, error) {
	if v, ok := staffsMap[id]; ok {
		return v, nil
	}
	return nil, repository.ErrStaffNotFound
}

func (r *staffRepositoryMap) Create(staff *entity.Staff) (*entity.Staff, error) {
	if _, ok := staffsMap[staff.ID]; ok {
		return nil, repository.ErrStaffExists
	}
	staffsMap[staff.ID] = staff
	return staff, nil
}

func (r *staffRepositoryMap) Update(staff *entity.Staff) (*entity.Staff, error) {
	if _, ok := staffsMap[staff.ID]; !ok {
		return nil, repository.ErrStaffNotFound
	}
	staffsMap[staff.ID] = staff
	return staff, nil
}

func (r *staffRepositoryMap) Delete(id int) error {
	if _, ok := staffsMap[id]; !ok {
		return repository.ErrStaffNotFound
	}
	delete(staffsMap, id)
	return nil
}
