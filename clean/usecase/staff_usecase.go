package usecase

import (
	"clean-architecture/controller/in"
	"clean-architecture/controller/out"
	"clean-architecture/entity"
	"clean-architecture/repository"
)

type StaffUsecase interface {
	List() ([]*out.StaffResponse, error)
	Get(id int) (*out.StaffResponse, error)
	Create(*in.StaffRequest) (*out.StaffResponse, error)
	Update(*in.StaffRequest) (*out.StaffResponse, error)
	Delete(*in.DeleteStaffRequest) error
}

type StaffUsecaseImpl struct {
	staffRepo repository.StaffRepository
}

func NewStaffUseCase() *StaffUsecaseImpl {
	return &StaffUsecaseImpl{
		staffRepo: repository.NewStaffRepositoryMap(),
	}
}

func (u *StaffUsecaseImpl) List() ([]*out.StaffResponse, error) {
	staffs, err := u.staffRepo.List()
	if err != nil {
		return nil, err
	}
	var staffResponses []*out.StaffResponse
	for _, staff := range staffs {
		staffResponses = append(staffResponses, &out.StaffResponse{
			ID:    staff.ID,
			Name:  staff.Name,
			Email: staff.Email,
		})
	}
	return staffResponses, nil
}

func (u *StaffUsecaseImpl) Get(id int) (*out.StaffResponse, error) {
	staff, err := u.staffRepo.Get(id)
	if err != nil {
		return nil, err
	}
	return &out.StaffResponse{
		ID:    staff.ID,
		Name:  staff.Name,
		Email: staff.Email,
	}, nil
}

func (u *StaffUsecaseImpl) Create(staff *in.StaffRequest) (*out.StaffResponse, error) {
	s := &entity.Staff{
		ID:    staff.ID,
		Name:  staff.Name,
		Email: staff.Email,
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	createdStaff, err := u.staffRepo.Create(s)
	if err != nil {
		return nil, err
	}
	return &out.StaffResponse{
		ID:    createdStaff.ID,
		Name:  createdStaff.Name,
		Email: createdStaff.Email,
	}, nil
}

func (u *StaffUsecaseImpl) Update(staff *in.StaffRequest) (*out.StaffResponse, error) {
	s := &entity.Staff{
		ID:    staff.ID,
		Name:  staff.Name,
		Email: staff.Email,
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	updatedStaff, err := u.staffRepo.Update(s)
	if err != nil {
		return nil, err
	}

	return &out.StaffResponse{
		ID:    updatedStaff.ID,
		Name:  updatedStaff.Name,
		Email: updatedStaff.Email,
	}, nil
}

func (u *StaffUsecaseImpl) Delete(staffID *in.DeleteStaffRequest) error {
	return u.staffRepo.Delete(staffID.ID)
}
