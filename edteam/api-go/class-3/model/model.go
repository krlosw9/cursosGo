package model

import "errors"

var (
	ErrPersonCanNotBeNil    = errors.New("la persona no puede ser null")
	ErrIDPersonDoesNotExist = errors.New("persona no existe")
)
