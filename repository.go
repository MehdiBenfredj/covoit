package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"log"

	"gorm.io/driver/postgres"
)

type Repository interface {
	GetAllUsers() []User
	GetUserByEmail(email string) User
	GetUserById(userID uuid.UUID) User
	CreateNewUser(user User) bool
	DeleteUser(userID uuid.UUID) bool
	UpdateUser(user User) User

	GetAllRides() ([]Ride, error)
	GetRideById(rideID uuid.UUID) (Ride, error)
	CreateRide(ride Ride) (Ride, error)
	DeleteRide(rideID uuid.UUID) error
	UpdateRide(ride Ride) (Ride, error)

	GetAllBookings() ([]Booking, error)
	GetBookingById(bookingID uuid.UUID) (Booking, error)
	CreateBooking(booking Booking) (Booking, error)
	DeleteBooking(bookingID uuid.UUID) error
	UpdateBooking(booking Booking) (Booking, error)
}

type CovoitRepository struct {
	db *gorm.DB
}

func NewCovoitRepository() *CovoitRepository {
	dsn := "host=localhost user=mehdi password=mehdi dbname=covoit port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Enable the uuid-ossp extension for UUID generation
	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	// Auto-migrate tables
	err = db.AutoMigrate(&User{}, &Ride{}, &Booking{})
	if err != nil {
		log.Fatal("Auto migration failed:", err)
	}

	fmt.Println("Tables created or already exist!")
	return &CovoitRepository{db: db}
}

func (repository *CovoitRepository) GetAllUsers() []User {
	users := []User{}
	repository.db.Find(&users)
	return users
}
func (repository *CovoitRepository) GetUserByEmail(email string) User {
	user := User{}
	// Get first matched record
	repository.db.Where("email = ?", email).First(&user)

	return user
}
func (repository *CovoitRepository) GetUserById(userID uuid.UUID) User {
	ctx := context.Background()

	// Using numeric primary key
	user, err := gorm.G[User](repository.db).Where("user_id = ?", userID).First(ctx)
	if err != nil {
		fmt.Errorf("could not find user: %s", err)
	}
	return user
}
func (repository *CovoitRepository) CreateNewUser(user User) bool {
	result := repository.db.Create(&user)
	return result.Error == nil
}
func (repository *CovoitRepository) DeleteUser(userID uuid.UUID) bool {
	ctx := context.Background()

	_, err := gorm.G[User](repository.db).Where("user_id = ?", userID).Delete(ctx)
	if err != nil {
		fmt.Errorf("could not delete user: %s", err)
	}
	return err == nil
}

// TODO
func (repository *CovoitRepository) UpdateUser(user User) User {
	return user
}
func (repository *CovoitRepository) GetAllRides() ([]Ride, error) {
	rides := []Ride{}
	repository.db.Find(&rides)
	return rides, nil
}
func (repository *CovoitRepository) GetRideById(rideID uuid.UUID) (Ride, error) {
	ctx := context.Background()
	ride, err := gorm.G[Ride](repository.db).Where("ride_id", rideID).First(ctx)
	if err != nil {
		return Ride{}, fmt.Errorf("Ride %v not found, err : %s", rideID, err)
	}
	return ride, nil
}
func (repository *CovoitRepository) CreateRide(ride Ride) (Ride, error) {
	ctx := context.Background()
	err := gorm.G[Ride](repository.db).Create(ctx, &ride)
	if err != nil {
		return Ride{}, fmt.Errorf("could not create ride, err : %s", err)
	}
	return ride, nil
}
func (repository *CovoitRepository) DeleteRide(rideID uuid.UUID) error {
	ctx := context.Background()
	_, err := gorm.G[Ride](repository.db).Where("ride_id = ?", rideID).Delete(ctx)
	return err
}
func (repository *CovoitRepository) UpdateRide(ride Ride) (Ride, error) {
	return Ride{}, nil
}
func (repository *CovoitRepository) GetAllBookings() ([]Booking, error) {
	ctx := context.Background()
	bookings, err := gorm.G[Booking](repository.db).Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not get bookings, err : %s", err)
	}
	return bookings, nil
}
func (repository *CovoitRepository) GetBookingById(bookingID uuid.UUID) (Booking, error) {
	ctx := context.Background()
	booking, err := gorm.G[Booking](repository.db).Where("booking_id = ?", bookingID).First(ctx)
	if err != nil {
		return Booking{}, nil
	}
	return booking, nil
}
func (repository *CovoitRepository) CreateBooking(booking Booking) (Booking, error) {
	ctx := context.Background()
	err := gorm.G[Booking](repository.db).Create(ctx, &booking)
	if err != nil {
		return Booking{}, fmt.Errorf("could not create booking %v, err : %s", booking, err)
	}
	return booking, err
}
func (repository *CovoitRepository) DeleteBooking(bookingID uuid.UUID) error {
	ctx := context.Background()
	_, err := gorm.G[Booking](repository.db).Where("booking_id = ?", bookingID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("could not delete booking %s, err : %s", bookingID, err)
	}
	return nil
}
func (repository *CovoitRepository) UpdateBooking(booking Booking) (Booking, error) {
	return Booking{}, nil
}
