-- Active: 1755527623566@@127.0.0.1@5432@covoit
-- Enable uuid extension
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Users table
CREATE TABLE IF NOT EXISTS users (
    user_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    phone TEXT,
    address TEXT
);

-- Rides table
CREATE TABLE IF NOT EXISTS rides (
    ride_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    origin TEXT NOT NULL,
    destination TEXT NOT NULL,
    driver_id UUID NOT NULL REFERENCES users(user_id),
    departure_time TIMESTAMP NOT NULL,
    arrival_time TIMESTAMP NOT NULL,
    distance FLOAT,
    price FLOAT,
    number_of_seats INT
);

-- Bookings table
CREATE TABLE IF NOT EXISTS bookings (
    booking_id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    ride_id UUID NOT NULL REFERENCES rides(ride_id),
    user_id UUID NOT NULL REFERENCES users(user_id),
    number_of_seats INT,
    total_price FLOAT,
    booking_time TIMESTAMP NOT NULL
);
