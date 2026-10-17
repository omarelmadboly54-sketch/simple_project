package apperrors

import "errors"

var (
	// User
	ErrUserAlreadyExists  = errors.New("user already exists")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrUserNotFound       = errors.New("user not found")

	// Movie
	ErrMovieNotFound = errors.New("movie not found")

	// Showtime
	ErrShowtimeNotFound  = errors.New("showtime not found")
	ErrHallAlreadyBooked = errors.New("hall is already booked at this time slot")

	// Seat
	ErrSeatAlreadyBooked = errors.New("seat is already booked")
	ErrSeatNotFound      = errors.New("seat not found")
)
