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
	GetAllUsers() ([]User, error)
	GetUserByEmail(email string) (User, error)
	GetUserById(userID uuid.UUID) (User, error)
	CreateNewUser(user User) (User, error)
	DeleteUser(userID uuid.UUID) error
	UpdateUser(user User) (User, error)

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

func (repository *CovoitRepository) GetAllUsers() ([]User, error) {
	ctx := context.Background()
	users, err := gorm.G[User](repository.db).Find(ctx)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve users, err : %s", err)
	}
	return users, nil
}

func (repository *CovoitRepository) GetUserByEmail(email string) (User, error) {
	ctx := context.Background()
	user, err := gorm.G[User](repository.db).Where("email = ?", email).First(ctx)
	if err != nil {
		return User{}, fmt.Errorf("could not retrieve user with email %s, err : %s", email, err)
	}
	return user, nil
}

func (repository *CovoitRepository) GetUserById(userID uuid.UUID) (User, error) {
	ctx := context.Background()
	user, err := gorm.G[User](repository.db).Where("user_id = ?", userID).First(ctx)
	if err != nil {
		return User{}, fmt.Errorf("could not retrieve user with id %s, err : %s", userID, err)
	}
	return user, nil
}
func (repository *CovoitRepository) CreateNewUser(user User) (User, error) {
	ctx := context.Background()
	err := gorm.G[User](repository.db).Create(ctx, &user)
	if err != nil {
		return User{}, fmt.Errorf("could not create user %v, err : %s", user, err)
	}
	return user, nil
}

func (repository *CovoitRepository) DeleteUser(userID uuid.UUID) error {
	ctx := context.Background()

	_, err := gorm.G[User](repository.db).Where("user_id = ?", userID).Delete(ctx)
	if err != nil {
		return fmt.Errorf("could not delete user %s, err : %s", userID, err)
	}
	return nil
}

// TODO
func (repository *CovoitRepository) UpdateUser(user User) (User, error) {
	return user, nil
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
		return Booking{}, fmt.Errorf("Booking %v not found, err : %s", bookingID, err)
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
