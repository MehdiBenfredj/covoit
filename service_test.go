package main

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestGetAllUsers(t *testing.T) {
	db := CreateNewMockDB(t)
	s := CovoitService{&MockRepository{db}}
	users, err := s.GetAllUsers()
	if err != nil {
		t.Errorf("could not retrieve all users, err : %s", err)
	}

	if len(users) < 1 {
		t.Errorf("there are no users in db")
	}
}
func TestUserService(t *testing.T) {
	db := CreateNewMockDB(t)
	s := CovoitService{&MockRepository{db}}
	t.Run("test get all users", func(t *testing.T) {
		users, err := s.GetAllUsers()
		if err != nil {
			t.Errorf("could not retrieve all users, err : %s", err)
		}

		if len(users) < 1 {
			t.Errorf("there are no users in db")
		}
	})
	t.Run("test get user by email", func(t *testing.T) {
		want := User{
			UserID:    stringToUuid(t, "652c99d0-39a5-4797-97a6-09eba33f2bd7"),
			FirstName: "Mehdi",
			LastName:  "BENFREDJ",
			Email:     "mehdibenfredj3@gmail.com",
		}
		got, err := s.GetUserByEmail("mehdibenfredj3@gmail.com")
		if err != nil {
			t.Errorf("could not get user %s, err %s", "mehdibenfredj3@gmail.com", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
	t.Run("test get user by id", func(t *testing.T) {
		want := User{
			UserID:    stringToUuid(t, "652c99d0-39a5-4797-97a6-09eba33f2bd7"),
			FirstName: "Mehdi",
			LastName:  "BENFREDJ",
			Email:     "mehdibenfredj3@gmail.com",
		}
		got, err := s.GetUserById(stringToUuid(t, "652c99d0-39a5-4797-97a6-09eba33f2bd7"))
		if err != nil {
			t.Errorf("could not get user %s, err %s", "652c99d0-39a5-4797-97a6-09eba33f2bd7", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
	t.Run("test create & delete user", func(t *testing.T) {
		u := User{
			FirstName: "Faten",
			LastName:  "Sayeh",
			Email:     "sayehfaten1195@gmail.com",
		}
		user, err := s.CreateNewUser(u)
		if err != nil || len(db.Users) != 3 {
			t.Errorf("could not create user %s, err : %s", "sayehfaten1195@gmail.com", err)
		}

		if !reflect.DeepEqual(user, u) {
			t.Errorf("created : %v, want : %v", user, u)
		}

		err = s.DeleteUser(stringToUuid(t, "652c99d0-39a5-4797-97a6-09eba33f2bd7"))
		if err != nil || len(db.Users) != 2 {
			fmt.Printf("%d", len(db.Users))
			t.Errorf("could not delete user %s, err : %s", "652c99d0-39a5-4797-97a6-09eba33f2bd7", err)

		}
	})
	t.Run("test update user", func(t *testing.T) {})
}

func TestRideService(t *testing.T) {
	db := CreateNewMockDB(t)
	s := CovoitService{&MockRepository{db}}

	t.Run("test get all rides", func(t *testing.T) {
		rides, err := s.GetAllRides()
		if err != nil {
			t.Errorf("could not retrieve all rides, err : %s", err)
		}

		if len(rides) < 1 {
			t.Errorf("there are no rides in db")
		}
	})
	t.Run("test get ride by id", func(t *testing.T) {
		want := Ride{
			RideID:      stringToUuid(t, "630cbfed-d023-41a4-884c-b1b1de76fb9f"),
			Origin:      "Constantine",
			Destination: "Alger",
		}
		got, err := s.GetRideById(stringToUuid(t, "630cbfed-d023-41a4-884c-b1b1de76fb9f"))
		if err != nil {
			t.Errorf("could not get ride %s, err %s", "630cbfed-d023-41a4-884c-b1b1de76fb9f", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
	t.Run("test create & delete ride", func(t *testing.T) {
		r := Ride{
			Origin:      "Constantine",
			Destination: "Alger",
		}
		ride, err := s.CreateRide(r)
		if err != nil || len(db.Rides) != 3 {
			t.Errorf("could not create ride %v, err : %s", r, err)
		}

		if !reflect.DeepEqual(ride, r) {
			t.Errorf("created : %v, want : %v", ride, r)
		}

		err = s.DeleteRide(stringToUuid(t, "ef5e1eda-e5e0-4f90-81ac-110b0bf84281"))
		if err != nil || len(db.Rides) != 2 {
			fmt.Printf("%d", len(db.Rides))
			t.Errorf("could not delete ride %s, err : %s", "ef5e1eda-e5e0-4f90-81ac-110b0bf84281", err)

		}
	})
	t.Run("test update ride", func(t *testing.T) {})
}

func TestBookingService(t *testing.T) {
	db := CreateNewMockDB(t)
	s := CovoitService{&MockRepository{db}}

	t.Run("test get all bookings", func(t *testing.T) {
		bookings, err := s.GetAllBookings()
		if err != nil {
			t.Errorf("could not retrieve all bookings, err : %s", err)
		}

		if len(bookings) < 1 {
			t.Errorf("there are no bookings in db")
		}
	})
	t.Run("test get booking by id", func(t *testing.T) {
		want := Booking{
			BookingID: stringToUuid(t, "ac925d60-1455-4d17-baeb-c4ffd4ed8205"),
			RideID:    stringToUuid(t, "46f45ea1-3f50-45cb-8556-797fe2688566"),
			UserID:    stringToUuid(t, "3c05d41e-344c-4661-a5fd-63e7a0a46998"),
		}
		got, err := s.GetBookingById(stringToUuid(t, "ac925d60-1455-4d17-baeb-c4ffd4ed8205"))
		if err != nil {
			t.Errorf("could not get booking %s, err %s", "ac925d60-1455-4d17-baeb-c4ffd4ed8205", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got : %v, want : %v", got, want)
		}
	})
	t.Run("test create & delete  booking", func(t *testing.T) {
		b := Booking{
			RideID: stringToUuid(t, "46f45ea1-3f50-45cb-8556-797fe2688566"),
			UserID: stringToUuid(t, "3c05d41e-344c-4661-a5fd-63e7a0a46998"),
		}
		booking, err := s.CreateBooking(b)
		if err != nil || len(db.Bookings) != 2 {
			t.Errorf("could not create booking %v, err : %s", b, err)
		}

		if !reflect.DeepEqual(booking, b) {
			t.Errorf("created : %v, want : %v", booking, b)
		}

		err = s.DeleteBooking(stringToUuid(t, "ac925d60-1455-4d17-baeb-c4ffd4ed8205"))
		if err != nil || len(db.Bookings) != 1 {
			fmt.Printf("%d", len(db.Bookings))
			t.Errorf("could not delete booking %s, err : %s", "ac925d60-1455-4d17-baeb-c4ffd4ed8205", err)

		}
	})
	t.Run("test update booking", func(t *testing.T) {})
}

type MockDB struct {
	Users    []User
	Bookings []Booking
	Rides    []Ride
}

type MockRepository struct {
	DB *MockDB
}

func CreateNewMockDB(t *testing.T) *MockDB {
	return &MockDB{
		Users: []User{
			User{
				UserID:    stringToUuid(t, "652c99d0-39a5-4797-97a6-09eba33f2bd7"),
				FirstName: "Mehdi",
				LastName:  "BENFREDJ",
				Email:     "mehdibenfredj3@gmail.com",
			},
			User{
				UserID:    stringToUuid(t, "90ed9f80-d22f-482a-8194-ec04cfeedcb2"),
				FirstName: "Faten",
				LastName:  "Sayeh",
				Email:     "sayehfaten1195@gmail.com",
			},
		},
		Rides: []Ride{
			Ride{
				RideID:      stringToUuid(t, "630cbfed-d023-41a4-884c-b1b1de76fb9f"),
				Origin:      "Constantine",
				Destination: "Alger"},
			Ride{
				RideID:      stringToUuid(t, "ef5e1eda-e5e0-4f90-81ac-110b0bf84281"),
				Origin:      "Marseille",
				Destination: "Paris"},
		},
		Bookings: []Booking{
			Booking{
				BookingID: stringToUuid(t, "ac925d60-1455-4d17-baeb-c4ffd4ed8205"),
				RideID:    stringToUuid(t, "46f45ea1-3f50-45cb-8556-797fe2688566"),
				UserID:    stringToUuid(t, "3c05d41e-344c-4661-a5fd-63e7a0a46998")}},
	}
}

func (m *MockRepository) GetAllUsers() ([]User, error) {
	return m.DB.Users, nil
}
func (m *MockRepository) GetUserByEmail(email string) (User, error) {
	for _, user := range m.DB.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found with email : %s", email)
}

func (m *MockRepository) GetUserById(userID uuid.UUID) (User, error) {
	for _, user := range m.DB.Users {
		if user.UserID == userID {
			return user, nil
		}
	}
	return User{}, fmt.Errorf("user not found with userID : %s", userID)
}

func (m *MockRepository) CreateNewUser(user User) (User, error) {
	m.DB.Users = append(m.DB.Users, user)
	return user, nil
}

func (m *MockRepository) DeleteUser(userID uuid.UUID) error {
	for i, user := range m.DB.Users {
		if user.UserID == userID {
			m.DB.Users = append(m.DB.Users[:i], m.DB.Users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not delete user : %s", userID)
}

func (m *MockRepository) UpdateUser(user User) (User, error) {
	return user, nil
}

func (m *MockRepository) GetAllRides() ([]Ride, error) {
	return m.DB.Rides, nil
}

func (m *MockRepository) GetRideById(rideID uuid.UUID) (Ride, error) {
	for _, ride := range m.DB.Rides {
		if ride.RideID == rideID {
			return ride, nil
		}
	}
	return Ride{}, fmt.Errorf("ride not found with userID : %s", rideID)
}

func (m *MockRepository) CreateRide(ride Ride) (Ride, error) {
	m.DB.Rides = append(m.DB.Rides, ride)
	return ride, nil
}

func (m *MockRepository) DeleteRide(rideID uuid.UUID) error {
	for i, ride := range m.DB.Rides {
		if ride.RideID == rideID {
			m.DB.Rides = append(m.DB.Rides[:i], m.DB.Rides[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not delete ride %s", rideID)
}

func (m *MockRepository) UpdateRide(ride Ride) (Ride, error) {
	return ride, nil
}

func (m *MockRepository) GetAllBookings() ([]Booking, error) {
	return m.DB.Bookings, nil
}

func (m *MockRepository) GetBookingById(bookingID uuid.UUID) (Booking, error) {
	for _, booking := range m.DB.Bookings {
		if booking.BookingID == bookingID {
			return booking, nil
		}
	}
	return Booking{}, fmt.Errorf("booking not found, booking id : %s ", bookingID)
}

func (m *MockRepository) CreateBooking(booking Booking) (Booking, error) {
	m.DB.Bookings = append(m.DB.Bookings, booking)
	return booking, nil
}

func (m *MockRepository) DeleteBooking(bookingID uuid.UUID) error {
	for i, booking := range m.DB.Bookings {
		if booking.BookingID == bookingID {
			m.DB.Bookings = append(m.DB.Bookings[:i], m.DB.Bookings[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("could not delete booking with id : %s", bookingID)
}

func (m *MockRepository) UpdateBooking(booking Booking) (Booking, error) {
	return booking, nil
}
