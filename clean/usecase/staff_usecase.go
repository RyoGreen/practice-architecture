package usecase

import (
	"clean-architecture/controller/in"
	"clean-architecture/controller/out"
	"clean-architecture/entity"
	"clean-architecture/infrastructure/postgres"
	"clean-architecture/repository"
)

type StaffUsecase interface {
	List() ([]*out.StaffOverviewResponse, error)
	Get(id int) (*out.StaffResponse, error)
	Create(*in.StaffRequest) (*out.StaffResponse, error)
	Update(*in.StaffRequest) (*out.StaffResponse, error)
	Delete(*in.DeleteStaffRequest) error
}

type StaffUsecaseImpl struct {
	staffRepo repository.StaffRepository
}

func NewStaffUseCase() StaffUsecase {
	return &StaffUsecaseImpl{
		staffRepo: postgres.NewStaffRepositoryPostgres(),
	}
}

type staffOverview struct {
	ID    int
	Name  string
	Email string
}

func (so *staffOverview) SomeFunction() {
	// something to do
}

func (u *StaffUsecaseImpl) List() ([]*out.StaffOverviewResponse, error) {
	staffs, err := u.staffRepo.List()
	if err != nil {
		return nil, err
	}

	var staffOverviews []*staffOverview
	for _, staff := range staffs {
		so := &staffOverview{
			ID:    staff.ID,
			Name:  staff.Name,
			Email: staff.Email,
		}
		so.SomeFunction()
		staffOverviews = append(staffOverviews, so)
	}

	var staffOverviewsReponse []*out.StaffOverviewResponse
	for _, v := range staffOverviews {
		staffOverviewsReponse = append(staffOverviewsReponse, &out.StaffOverviewResponse{
			ID:    v.ID,
			Name:  v.Name,
			Email: v.Email,
		})
	}

	return staffOverviewsReponse, nil
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
	s, err := entity.NewStaff(staff.ID, staff.Salary, staff.Name, staff.Email, staff.Address)
	if err != nil {
		return nil, err
	}

	createdStaff, err := u.staffRepo.Create(s)
	if err != nil {
		return nil, err
	}

	return &out.StaffResponse{
		ID:      createdStaff.ID,
		Name:    createdStaff.Name,
		Email:   createdStaff.Email,
		Salary:  createdStaff.Salary,
		Address: createdStaff.Adddress,
	}, nil
}

func (u *StaffUsecaseImpl) Update(staff *in.StaffRequest) (*out.StaffResponse, error) {
	s, err := entity.NewStaff(staff.ID, staff.Salary, staff.Name, staff.Email, staff.Address)
	if err != nil {
		return nil, err
	}

	updatedStaff, err := u.staffRepo.Update(s)
	if err != nil {
		return nil, err
	}

	return &out.StaffResponse{
		ID:      updatedStaff.ID,
		Name:    updatedStaff.Name,
		Email:   updatedStaff.Email,
		Salary:  updatedStaff.Salary,
		Address: updatedStaff.Adddress,
	}, nil
}

func (u *StaffUsecaseImpl) Delete(staffID *in.DeleteStaffRequest) error {
	return u.staffRepo.Delete(staffID.ID)
}
