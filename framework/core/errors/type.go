package errors

import (
	"errors"
)

type ErrorsInterface interface {
	New() func(message string) error
}

type pkgError struct{}

func (pkgError) New() func(message string) error {
	return errors.New
}

// ErrNotSuppportedFunctionality returns an error cause the functionality is not supported
var ErrNotSuppportedFunctionality = errors.New("not supported functionality")

// ErrNotImplementedFunctionality returns an error cause the functionality is not implemented
var ErrNotImplementedFunctionality = errors.New("not implemented functionality")

// ErrDuplicateDataExisted returns an error cause the functionality is not supported
var ErrDuplicateDataExisted = errors.New("duplicated data existed")
