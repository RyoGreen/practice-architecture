package repository

import (
	"errors"
)

var (
	ErrStaffNotFound = errors.New("staff not found")
	ErrStaffExists   = errors.New("staff already exists")
)
