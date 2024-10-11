package auth

import "errors"

var (
	ErrUserNotFound     = errors.New("user not found")
	ErrTokenExpired     = errors.New("token has expired")
	ErrTokenNotValidYet = errors.New("token not valid yet")
	ErrTokenInvalid     = errors.New("token is invalid")
	ErrInvalidPassword  = errors.New("invalid password")
	ErrInvalidRole      = errors.New("invalid role: user_type should be client or moderator")
	ErrGenerateToken    = errors.New("failed to generate token")
	ErrWriteResponse    = errors.New("failed to write response")
	ErrCreateUser       = errors.New("failed to create user")
	ErrLogin            = errors.New("failed to login")
)
