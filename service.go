package main

import (
	"github.com/google/uuid"
)

type Service interface {
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

type CovoitService struct {
	repository Repository
}

func (service *CovoitService) GetAllUsers() ([]User, error) {
	return service.repository.GetAllUsers()
}
func (service *CovoitService) GetUserByEmail(email string) (User, error) {
	return service.repository.GetUserByEmail(email)
}
func (service *CovoitService) GetUserById(userID uuid.UUID) (User, error) {
	return service.repository.GetUserById(userID)
}
func (service *CovoitService) CreateNewUser(user User) (User, error) {
	return service.repository.CreateNewUser(user)
}
func (service *CovoitService) DeleteUser(userID uuid.UUID) error {
	return service.repository.DeleteUser(userID)
}
func (service *CovoitService) UpdateUser(user User) (User, error) {
	return service.repository.UpdateUser(user)
}
func (service *CovoitService) GetAllRides() ([]Ride, error) {
	return service.repository.GetAllRides()
}
func (service *CovoitService) GetRideById(rideID uuid.UUID) (Ride, error) {
	return service.repository.GetRideById(rideID)
}
func (service *CovoitService) CreateRide(ride Ride) (Ride, error) {
	return service.repository.CreateRide(ride)
}
func (service *CovoitService) DeleteRide(rideID uuid.UUID) error {
	return service.repository.DeleteRide(rideID)
}
func (service *CovoitService) UpdateRide(ride Ride) (Ride, error) {
	return service.repository.UpdateRide(ride)
}
func (service *CovoitService) GetAllBookings() ([]Booking, error) {
	return service.repository.GetAllBookings()
}
func (service *CovoitService) GetBookingById(bookingID uuid.UUID) (Booking, error) {
	return service.repository.GetBookingById(bookingID)
}
func (service *CovoitService) CreateBooking(booking Booking) (Booking, error) {
	return service.repository.CreateBooking(booking)
}
func (service *CovoitService) DeleteBooking(bookingID uuid.UUID) error {
	return service.repository.DeleteBooking(bookingID)
}
func (service *CovoitService) UpdateBooking(booking Booking) (Booking, error) {
	return service.repository.UpdateBooking(booking)
}
