package main

import (
	"github.com/google/uuid"
)

type Service interface {
	GetAllUsers() []User
	GetUserByEmail(email string) User
	GetUserById(userID uuid.UUID) User
	CreateNewUser(user User) bool
	DeleteUser(userID uuid.UUID) bool
	UpdateUser(user User) User

	GetAllRides() []Ride
	GetRideById(rideID uuid.UUID) Ride
	CreateRide(ride Ride) bool
	DeleteRide(rideID uuid.UUID) bool
	UpdateRide(ride Ride) Ride

	GetAllBookings() []Booking
	GetBookingById(bookingID uuid.UUID) Booking
	CreateBooking(booking Booking) bool
	DeleteBooking(bookingID uuid.UUID) bool
	UpdateBooking(booking Booking) Booking
}

type CovoitService struct {
	repository Repository
}

func (service *CovoitService) GetAllUsers() []User {
	return service.repository.GetAllUsers()
}
func (service *CovoitService) GetUserByEmail(email string) User {
	return service.repository.GetUserByEmail(email)
}
func (service *CovoitService) GetUserById(userID uuid.UUID) User {
	return service.repository.GetUserById(userID)
}
func (service *CovoitService) CreateNewUser(user User) bool {
	return service.repository.CreateNewUser(user)
}
func (service *CovoitService) DeleteUser(userID uuid.UUID) bool {
	return service.repository.DeleteUser(userID)
}
func (service *CovoitService) UpdateUser(user User) User {
	return service.repository.UpdateUser(user)
}
func (service *CovoitService) GetAllRides() []Ride {
	return service.repository.GetAllRides()
}
func (service *CovoitService) GetRideById(rideID uuid.UUID) Ride {
	return service.repository.GetRideById(rideID)
}
func (service *CovoitService) CreateRide(ride Ride) bool {
	return service.repository.CreateRide(ride)
}
func (service *CovoitService) DeleteRide(rideID uuid.UUID) bool {
	return service.repository.DeleteRide(rideID)
}
func (service *CovoitService) UpdateRide(ride Ride) Ride {
	return service.repository.UpdateRide(ride)
}
func (service *CovoitService) GetAllBookings() []Booking {
	return service.repository.GetAllBookings()
}
func (service *CovoitService) GetBookingById(bookingID uuid.UUID) Booking {
	return service.repository.GetBookingById(bookingID)
}
func (service *CovoitService) CreateBooking(booking Booking) bool {
	return service.repository.CreateBooking(booking)
}
func (service *CovoitService) DeleteBooking(bookingID uuid.UUID) bool {
	return service.repository.DeleteBooking(bookingID)
}
func (service *CovoitService) UpdateBooking(booking Booking) Booking {
	return service.repository.UpdateBooking(booking)
}
