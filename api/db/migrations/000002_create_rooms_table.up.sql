CREATE TABLE IF NOT EXISTS rooms (
    id BIGINT PRIMARY KEY,
    listing_url VARCHAR(255) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    neighborhood_overview TEXT,
    picture_url VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2),
    bedrooms INT,
    beds INT,
    room_type VARCHAR(50) NOT NULL,
    property_type VARCHAR(50) NOT NULL,
    neighbourhood VARCHAR(100),
);