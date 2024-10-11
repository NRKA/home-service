package house

import "github.com/pkg/errors"

var (
	ErrInvalidRequest     = errors.New("invalid request")
	ErrCreateHouse        = errors.New("failed to create house")
	ErrInvalidID          = errors.New("id must be integer")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
