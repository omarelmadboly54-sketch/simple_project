CREATE TABLE IF NOT EXISTS show_times(
    id BIGSERIAL PRIMARY KEY,
    movie_id BIGINT NOT NULL ,
    show_time TIMESTAMP NOT NULL,
    price  INT NOT NULL,
    hall_number INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    FOREIGN KEY(movie_id) REFERENCES movies(id) ON DELETE CASCADE
);

