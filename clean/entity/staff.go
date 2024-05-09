package entity

import (
	"fmt"
	"net/mail"
)

type Staff struct {
	ID       int
	Name     string
	Email    string
	Salary   int
	Adddress string
}

var (
	ErrStaffValidate = fmt.Errorf("staff name, address or email is empty")
	ErrStaffEmail    = fmt.Errorf("staff email is invalid")
	ErrStaffSalary   = fmt.Errorf("staff salary is invalid")
)

func NewStaff(id, salary int, name, email, address string) (*Staff, error) {
	s := &Staff{
		ID:       id,
		Name:     name,
		Email:    email,
		Adddress: address,
		Salary:   salary,
	}
	if err := s.Validate(); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Staff) Validate() error {
	if s.Name == "" || s.Email == "" || s.Adddress == "" {
		return ErrStaffValidate
	}
	if s.Salary < 0 {
		return ErrStaffSalary
	}
	if _, err := mail.ParseAddress(s.Email); err != nil {
		return ErrStaffEmail
	}
	return nil
}
