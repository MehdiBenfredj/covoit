package main

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	UserID    uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	FirstName string
	LastName  string
	Email     string `gorm:"uniqueIndex"`
	Phone     string
	Address   string
	Bookings  []Booking `gorm:"foreignKey:UserID"`
}

type Ride struct {
	RideID        uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	Origin        string
	Destination   string
	DriverID      uuid.UUID
	DepartureTime time.Time
	ArrivalTime   time.Time
	Distance      float64
	Price         float64
	NumberOfSeats int
	Bookings      []Booking `gorm:"foreignKey:RideID"`
}

type Booking struct {
	BookingID     uuid.UUID `gorm:"type:uuid;primaryKey;default:uuid_generate_v4()"`
	RideID        uuid.UUID
	UserID        uuid.UUID
	NumberOfSeats int
	TotalPrice    float64
	BookingTime   time.Time
}
