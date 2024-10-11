package sender

import "errors"

var (
	ErrHouseNotFound = errors.New("house with the given ID does not exist")
	ErrInvalidEmail  = errors.New("invalid email")
	ErrInvalidID     = errors.New("invalid ID, ID must be integer")
)
