package entity

import (
	"fmt"
	"net/mail"
)

type Staff struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type StaffRepository interface {
	List() ([]*Staff, error)
	Get(id int) (*Staff, error)
	Create(*Staff) (*Staff, error)
	Update(*Staff) (*Staff, error)
	Delete(id int) error
}

var (
	ErrStaffValidate = fmt.Errorf("staff name or email is empty")
	ErrStaffEmail    = fmt.Errorf("staff email is invalid")
)

func (s *Staff) Validate() error {
	if s.Name == "" || s.Email == "" {
		return ErrStaffValidate
	}
	_, err := mail.ParseAddress(s.Email)
	if err != nil {
		return ErrStaffEmail
	}
	return nil
}
