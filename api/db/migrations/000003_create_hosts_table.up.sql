CREATE TABLE IF NOT EXISTS hosts (
    id BIGINT PRIMARY KEY,
    host_url VARCHAR(200) NOT NULL,
    host_name VARCHAR(50) NOT NULL,
    host_since DATE,
    host_location VARCHAR(50),
    host_about TEXT,
    host_thumbnail_url VARCHAR(200),
    host_picture_url VARCHAR(200)
);