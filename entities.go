package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"  json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `gorm:"uniqueIndex" json:"email"`
	Phone     string    `json:"phone"`
	Address   string    `json:"adress"`
	Bookings  []Booking `gorm:"foreignKey:UserID" json:"bookings"`
}

type Ride struct {
	RideID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"ride_id"`
	Origin        string    `json:"origin"`
	Destination   string    `json:"destination"`
	DriverID      uuid.UUID `json:"driver_id"`
	DepartureTime time.Time `json:"departure_time"`
	ArrivalTime   time.Time `json:"arrival_time"`
	Distance      float64   `json:"distance"`
	Price         float64   `json:"price"`
	NumberOfSeats int       `json:"number_of_seats"`
	Bookings      []Booking `gorm:"foreignKey:RideID" json:"bookings"`
}

type Booking struct {
	BookingID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()" json:"booking_id"`
	RideID        uuid.UUID `json:"ride_id"`
	UserID        uuid.UUID `json:"user_id"`
	NumberOfSeats int       `json:"number_of_seats"`
	TotalPrice    float64   `json:"total_price"`
	BookingTime   time.Time `json:"booking_time"`
}
