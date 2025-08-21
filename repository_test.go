package main

import (
	"context"
	"slices"
	"testing"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func TestNewCovoitRepository(t *testing.T) {
	repository := NewCovoitRepository()
	want := []string{"users", "bookings", "rides"}
	ctx := context.Background()
	got, err := gorm.G[string](repository.db).Raw(`SELECT tablename FROM pg_catalog.pg_tables
													WHERE schemaname != 'pg_catalog' AND 
    												schemaname != 'information_schema';`).Find(ctx)
	if err != nil {
		t.Errorf("Could not get db tables")
	}

	if !slices.Equal(got, want) {
		t.Errorf("want %s, got %s", want, got)
	}
}

func TestUserRepo(t *testing.T) {
	repository := NewCovoitRepository()
	t.Run("Test get all users", func(t *testing.T) {
		users, err := repository.GetAllUsers()
		if err != nil {
			t.Errorf("Could not get all users")
		}

		if len(users) < 1 {
			t.Errorf("could not retrieve any user from users table")
		}
	})

	t.Run("Test get user by email", func(t *testing.T) {
		want := User{
			FirstName: "Mehdi",
			LastName:  "BENFREDJ",
		}
		got, err := repository.GetUserByEmail("mehdibenfredj3@gmail.com")
		if err != nil {
			t.Errorf("Could get user with email : %s", "mehdibenfredj3@gmail.com")
		}

		if want.FirstName != got.FirstName && want.LastName != got.LastName {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Test get user by id", func(t *testing.T) {
		want := User{
			FirstName: "Mehdi",
			LastName:  "BENFREDJ",
		}
		userID, err := uuid.Parse("3c05d41e-344c-4661-a5fd-63e7a0a46998")
		if err != nil {
			t.Errorf("Could not parse uuid : %s", "3c05d41e-344c-4661-a5fd-63e7a0a46998")
		}
		got, err := repository.GetUserById(userID)
		if err != nil {
			t.Errorf("Could get user with id : %v", userID)
		}

		if want.FirstName != got.FirstName && want.LastName != got.LastName {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Test creating && deleting user", func(t *testing.T) {
		want := User{FirstName: "Lionel", LastName: "Messi", Phone: "10", Address: "Buenos Aires, Argentina", Email: "lionel.messi@fcbarcelona.es"}
		got, err := repository.CreateNewUser(User{FirstName: "Lionel", LastName: "Messi", Phone: "10", Address: "Buenos Aires, Argentina", Email: "lionel.messi@fcbarcelona.es"})
		if err != nil {
			t.Errorf("could not insert the goat 🐐")
		}

		if got.FirstName != want.FirstName &&
			got.LastName != want.LastName &&
			got.Phone != want.Phone &&
			got.Address != want.Address &&
			got.Email != want.Email {
			t.Errorf("want %v, got %v", want, got)
		}

		err = repository.DeleteUser(got.UserID)
		if err != nil {
			t.Errorf("could not delete the goat 🐐")
		}

		got, err = repository.GetUserByEmail(want.Email)
		if err == nil || got.Email == want.Email {
			t.Errorf("goat 🐐 still in db, probably not deleted, err : %s", err)
		}
	})
}

func TestRideRepo(t *testing.T) {
	repository := NewCovoitRepository()
	t.Run("Test get all rides", func(t *testing.T) {
		rides, err := repository.GetAllRides()
		if err != nil {
			t.Errorf("Could not get all rides")
		}

		if len(rides) < 1 {
			t.Errorf("could not retrieve any ride from rides table")
		}
	})

	t.Run("Test get ride by id", func(t *testing.T) {
		want := Ride{
			Origin:      "Constantine",
			Destination: "Alger",
		}
		rideID := stringToUuid(t, "f71533c2-874a-459e-806e-ae2200351c84")
		got, err := repository.GetRideById(rideID)
		if err != nil {
			t.Errorf("Could get ride with id : %v", rideID)
		}

		if want.Origin != got.Origin && want.Destination != got.Destination {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Test creating && deleting ride", func(t *testing.T) {
		want :=
			Ride{Origin: "Tlemcen",
				Destination:   "Tebessa",
				DepartureTime: time.Date(2025, 03, 24, 22, 34, 0, 0, time.UTC),
				ArrivalTime:   time.Date(2025, 03, 25, 22, 34, 0, 0, time.UTC),
				Distance:      1200,
				Price:         160,
				NumberOfSeats: 4}
		got, err := repository.CreateRide(
			Ride{Origin: "Tlemcen",
				Destination:   "Tebessa",
				DepartureTime: time.Date(2025, 03, 24, 22, 34, 0, 0, time.UTC),
				ArrivalTime:   time.Date(2025, 03, 25, 22, 34, 0, 0, time.UTC),
				Distance:      1200,
				Price:         160,
				NumberOfSeats: 4})

		if err != nil {
			t.Errorf("could not insert the ride")
		}

		if got.Destination != want.Destination &&
			got.Origin != want.Origin &&
			got.DepartureTime != want.DepartureTime &&
			got.ArrivalTime != want.ArrivalTime &&
			got.Price != want.Price &&
			got.NumberOfSeats != want.NumberOfSeats &&
			got.Distance != want.Distance {
			t.Errorf("want %v, got %v", want, got)
		}

		err = repository.DeleteRide(got.RideID)
		if err != nil {
			t.Errorf("could not delete the ride")
		}

		got, err = repository.GetRideById(want.RideID)
		if err == nil || got.Origin == want.Origin {
			t.Errorf("ride still in db, probably not deleted, err : %s", err)
		}
	})
}

func TestBookingRepo(t *testing.T) {
	repository := NewCovoitRepository()
	t.Run("Test get all bookings", func(t *testing.T) {
		bookings, err := repository.GetAllBookings()
		if err != nil {
			t.Errorf("Could not get all bookings")
		}

		if len(bookings) < 1 {
			t.Errorf("could not retrieve any booking from bookings table")
		}
	})

	t.Run("Test get booking by id", func(t *testing.T) {
		want := Booking{
			RideID: stringToUuid(t, "46f45ea1-3f50-45cb-8556-797fe2688566"),
			UserID: stringToUuid(t, "3c05d41e-344c-4661-a5fd-63e7a0a46998"),
		}
		bookingID := stringToUuid(t, "5a0977da-ac5d-4daf-8b91-4bfcf8e9866a")
		got, err := repository.GetBookingById(bookingID)
		if err != nil {
			t.Errorf("Could get booking with id : %v", bookingID)
		}

		if want.RideID != got.RideID && want.UserID != got.UserID {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Test creating && deleting booking", func(t *testing.T) {
		want := Booking{
			RideID: stringToUuid(t, "46f45ea1-3f50-45cb-8556-797fe2688566"),
			UserID: stringToUuid(t, "3c05d41e-344c-4661-a5fd-63e7a0a46998"),
		}
		got, err := repository.CreateBooking(want)

		if err != nil {
			t.Errorf("could not insert the booking")
		}

		if got.RideID != want.RideID &&
			got.UserID != want.UserID &&
			got.NumberOfSeats != want.NumberOfSeats &&
			got.TotalPrice != want.TotalPrice &&
			got.BookingTime != want.BookingTime {
			t.Errorf("want %v, got %v", want, got)
		}

		err = repository.DeleteBooking(got.BookingID)
		if err != nil {
			t.Errorf("could not delete the booking")
		}

		got, err = repository.GetBookingById(got.BookingID)
		if err == nil {
			t.Errorf("booking still in db, probably not deleted, err : %s", err)
		}
	})
}

func stringToUuid(t *testing.T, id string) uuid.UUID {
	t.Helper()
	res, err := uuid.Parse(id)
	if err != nil {
		t.Errorf("Could not parse uuid : %s", id)
	}
	return res
}
