Showtime & Seat Booking System:

A robust, production-grade backend API built in Go for managing showtimes, venues, and atomic seat reservations. Designed with high concurrency in mind to prevent double-booking and race conditions during high-traffic events.

Key Features:

Atomic Seat Bookings: Implements database-level locking and race-condition safety so that concurrent requests for the exact same seat result in exactly 1 success (200 OK) and multiple conflicts (409 Conflict).

JWT Authentication: Secure user registration, login, and token-based authorization middleware.

Concurrency Testing Suite: Includes dedicated Go test files leveraging synchronization barriers (sync.WaitGroup and channels) to simulate simultaneous multi-user booking attempts.

RESTful Architecture: Clean separation of concerns with structured routing (using Go Gin or standard library) and PostgreSQL database integration.

Tech Stack

Language: Go (Golang)

Database: PostgreSQL

Routing & Middleware: Gin Gonic / Native HTTP

Authentication: JSON Web Tokens (JWT)

Testing: Go testing package with concurrency primitives

Project structure:
└── omarelmadboly54-sketch-simple_project/
    ├── go.mod
    ├── go.sum
    ├── justfile
    ├── main.go
    ├── seat_test.go
    ├── configuration/
    │   └── config.go
    ├── internal/
    │   ├── controller/
    │   │   ├── movies.go
    │   │   ├── showtime_seats.go
    │   │   ├── showtimes.go
    │   │   └── users.go
    │   ├── middleware/
    │   │   └── middleware.go
    │   ├── models/
    │   │   ├── movies.go
    │   │   ├── showtime_seats.go
    │   │   ├── showtimes.go
    │   │   └── users.go
    │   ├── presenters/
    │   │   ├── movies.go
    │   │   ├── showtimes.go
    │   │   ├── showtimeSeat.go
    │   │   └── users.go
    │   ├── providers/
    │   │   └── postgres/
    │   │       ├── movies.go
    │   │       ├── showtime_seats.go
    │   │       ├── showtimes.go
    │   │       └── users.go
    │   ├── repository/
    │   │   ├── movies.go
    │   │   ├── showtime_seats.go
    │   │   ├── showtimes.go
    │   │   └── users.go
    │   ├── service/
    │   │   ├── movies.go
    │   │   ├── showtime_seats.go
    │   │   ├── showtimes.go
    │   │   └── users.go
    │   └── util/
    │       ├── password.go
    │       └── viper.go
    ├── migrations/
    │   ├── 000001_users.down.sql
    │   ├── 000001_users.up.sql
    │   ├── 000002_add_role_check.down.sql
    │   ├── 000002_add_role_check.up.sql
    │   ├── 000003_add_updated_at_column_in_user_table.down.sql
    │   ├── 000003_add_updated_at_column_in_user_table.up.sql
    │   ├── 000004_movies.down.sql
    │   ├── 000004_movies.up.sql
    │   ├── 000005_showtime_table.down.sql
    │   ├── 000005_showtime_table.up.sql
    │   ├── 000006_showtime_seats.down.sql
    │   └── 000006_showtime_seats.up.sql
    └── token/
        ├── jwt_maker.go
        ├── maker.go
        └── payload.go
URL=https://github.com/omarelmadboly54-sketch/simple_project
