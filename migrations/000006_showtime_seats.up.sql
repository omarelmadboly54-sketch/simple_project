CREATE TABLE IF NOT EXISTS showtime_seats(
    id  BIGSERIAL PRIMARY KEY,
    showtime_id BIGINT NOT NULL,
    seat_number VARCHAR(20) NOT NULL,
    status VARCHAR(20) NOT NULL,
    user_id BIGINT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    Foreign Key (showtime_id) REFERENCES show_times(id) ON DELETE CASCADE,
    Foreign Key (user_id) REFERENCES users (id) ON DELETE CASCADE
);
