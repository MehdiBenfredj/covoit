package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) GetUserByEmail(email string) (User, error) {
	args := m.Called(email)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockService) GetUserById(id uuid.UUID) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockService) GetAllUsers() ([]User, error) {
	args := m.Called()
	return args.Get(0).([]User), args.Error(1)
}

func (m *MockService) CreateNewUser(u User) (User, error) {
	args := m.Called(u)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockService) DeleteUser(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockService) GetAllRides() ([]Ride, error) {
	args := m.Called()
	return args.Get(0).([]Ride), args.Error(1)
}

func (m *MockService) CreateRide(r Ride) (Ride, error) {
	args := m.Called(r)
	return args.Get(0).(Ride), args.Error(1)
}

func (m *MockService) DeleteBooking(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockService) DeleteRide(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockService) GetAllBookings() ([]Booking, error) {
	args := m.Called()
	return args.Get(0).([]Booking), args.Error(1)
}

func (m *MockService) CreateBooking(b Booking) (Booking, error) {
	args := m.Called(b)
	return args.Get(0).(Booking), args.Error(1)
}

func (m *MockService) GetBookingById(id uuid.UUID) (Booking, error) {
	args := m.Called(id)
	return args.Get(0).(Booking), args.Error(1)
}

func (m *MockService) GetRideById(id uuid.UUID) (Ride, error) {
	args := m.Called(id)
	return args.Get(0).(Ride), args.Error(1)
}

func (m *MockService) UpdateBooking(b Booking) (Booking, error) {
	args := m.Called(b)
	return args.Get(0).(Booking), args.Error(1)
}

func (m *MockService) UpdateRide(r Ride) (Ride, error) {
	args := m.Called(r)
	return args.Get(0).(Ride), args.Error(1)
}

func (m *MockService) UpdateUser(u User) (User, error) {
	args := m.Called(u)
	return args.Get(0).(User), args.Error(1)
}

// -------- Tests --------

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	helloHandler(w, req)
	require.Equal(t, http.StatusOK, w.Result().StatusCode)
}

// ---- UsersHandler ----

func TestUsersHandler_GetByEmail(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}

	user := User{Email: "a@test.com"}
	mockSvc.On("GetUserByEmail", "a@test.com").Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/users?email=a@test.com", nil)
	w := httptest.NewRecorder()
	h.UsersHandler(w, req)

	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	// not found
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("GetUserByEmail", "notfound@test.com").Return(User{}, errors.New("not found"))
	req = httptest.NewRequest(http.MethodGet, "/users?email=notfound@test.com", nil)
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestUsersHandler_GetByID(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	uid := uuid.New()

	user := User{Email: "id@test.com"}
	mockSvc.On("GetUserById", uid).Return(user, nil)

	req := httptest.NewRequest(http.MethodGet, "/users?user_id="+uid.String(), nil)
	w := httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	// invalid UUID
	req = httptest.NewRequest(http.MethodGet, "/users?user_id=bad_uid", nil)
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// not found
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("GetUserById", uid).Return(User{}, errors.New("not found"))
	req = httptest.NewRequest(http.MethodGet, "/users?user_id="+uid.String(), nil)
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestUsersHandler_GetAll(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	users := []User{{Email: "u1"}, {Email: "u2"}}
	mockSvc.On("GetAllUsers").Return(users, nil)

	req := httptest.NewRequest(http.MethodGet, "/users", nil)
	w := httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("GetAllUsers").Return([]User{}, errors.New("db error"))
	req = httptest.NewRequest(http.MethodGet, "/users", nil)
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestUsersHandler_Post(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	user := User{Email: "post@test.com"}
	mockSvc.On("CreateNewUser", user).Return(user, nil)

	body, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	// bad JSON
	req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer([]byte("bad")))
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("CreateNewUser", user).Return(User{}, errors.New("fail"))
	req = httptest.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestUsersHandler_Delete(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	uid := uuid.New()

	mockSvc.On("DeleteUser", uid).Return(nil)
	req := httptest.NewRequest(http.MethodDelete, "/users?user_id="+uid.String(), nil)
	w := httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusNoContent, w.Result().StatusCode)

	// invalid UUID
	req = httptest.NewRequest(http.MethodDelete, "/users?user_id=bad", nil)
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("DeleteUser", uid).Return(errors.New("fail"))
	req = httptest.NewRequest(http.MethodDelete, "/users?user_id="+uid.String(), nil)
	w = httptest.NewRecorder()
	h.UsersHandler(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

// ---- RidesHandler ----

func TestRidesHandler_Get(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	rides := []Ride{{RideID: uuid.New()}}
	mockSvc.On("GetAllRides").Return(rides, nil)

	req := httptest.NewRequest(http.MethodGet, "/rides", nil)
	w := httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("GetAllRides").Return([]Ride{}, errors.New("fail"))
	req = httptest.NewRequest(http.MethodGet, "/rides", nil)
	w = httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestRidesHandler_Post(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	ride := Ride{RideID: uuid.New()}
	mockSvc.On("CreateRide", ride).Return(ride, nil)

	body, _ := json.Marshal(ride)
	req := httptest.NewRequest(http.MethodPost, "/rides", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusCreated, w.Result().StatusCode)

	// bad JSON
	req = httptest.NewRequest(http.MethodPost, "/rides", bytes.NewBuffer([]byte("bad")))
	w = httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("CreateRide", ride).Return(Ride{}, errors.New("fail"))
	req = httptest.NewRequest(http.MethodPost, "/rides", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestRidesHandler_Delete(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	uid := uuid.New()
	mockSvc.On("DeleteBooking", uid).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/rides?ride_id="+uid.String(), nil)
	w := httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusNoContent, w.Result().StatusCode)

	// invalid UUID
	req = httptest.NewRequest(http.MethodDelete, "/rides?ride_id=bad", nil)
	w = httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("DeleteBooking", uid).Return(errors.New("fail"))
	req = httptest.NewRequest(http.MethodDelete, "/rides?ride_id="+uid.String(), nil)
	w = httptest.NewRecorder()
	h.RidesHandler(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

// ---- BookingsHandler ----

func TestBookingsHandler_Get(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	bookings := []Booking{{BookingID: uuid.New()}}
	mockSvc.On("GetAllBookings").Return(bookings, nil)

	req := httptest.NewRequest(http.MethodGet, "/bookings", nil)
	w := httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusOK, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("GetAllBookings").Return([]Booking{}, errors.New("fail"))
	req = httptest.NewRequest(http.MethodGet, "/bookings", nil)
	w = httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusNotFound, w.Result().StatusCode)
}

func TestBookingsHandler_Post(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	booking := Booking{BookingID: uuid.New()}
	mockSvc.On("CreateBooking", booking).Return(booking, nil)

	body, _ := json.Marshal(booking)
	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusCreated, w.Result().StatusCode)

	// bad JSON
	req = httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer([]byte("bad")))
	w = httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("CreateBooking", booking).Return(Booking{}, errors.New("fail"))
	req = httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}

func TestBookingsHandler_Delete(t *testing.T) {
	mockSvc := new(MockService)
	h := &Handler{Service: mockSvc}
	uid := uuid.New()
	mockSvc.On("DeleteBooking", uid).Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/bookings?booking_id="+uid.String(), nil)
	w := httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusNoContent, w.Result().StatusCode)

	// invalid UUID
	req = httptest.NewRequest(http.MethodDelete, "/bookings?booking_id=bad", nil)
	w = httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusBadRequest, w.Result().StatusCode)

	// error
	mockSvc = new(MockService)
	h = &Handler{Service: mockSvc}
	mockSvc.On("DeleteBooking", uid).Return(errors.New("fail"))
	req = httptest.NewRequest(http.MethodDelete, "/bookings?booking_id="+uid.String(), nil)
	w = httptest.NewRecorder()
	h.BookingsHandler(w, req)
	require.Equal(t, http.StatusInternalServerError, w.Result().StatusCode)
}
