package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

type Handler struct {
	Service Service
}

func NewHandler() *Handler {
	repository := NewCovoitRepository()
	return &Handler{Service: &CovoitService{repository: repository}}
}

func (h *Handler) UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			if email := r.URL.Query().Get("email"); email != "" {
				user := h.Service.GetUserByEmail(email)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
				w.WriteHeader(http.StatusOK)
				return
			} else if idStr := r.URL.Query().Get("user_id"); idStr != "" {
				userID, err := uuid.Parse(idStr)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}
				user := h.Service.GetUserById(userID)
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
				w.WriteHeader(http.StatusOK)
				return
			}
			users := h.Service.GetAllUsers()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(users)
			w.WriteHeader(http.StatusOK)

		}

	case http.MethodPost:
		{
			newUser := User{}
			err := json.NewDecoder(r.Body).Decode(&newUser)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			done := h.Service.CreateNewUser(newUser)
			if !done {
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(newUser)
		}
	case http.MethodPatch:
		{

		}
	case http.MethodDelete:
		{
			idStr := r.URL.Query().Get("user_id")
			userID, err := uuid.Parse(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			done := h.Service.DeleteUser(userID)
			if !done {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}

}

func (h *Handler) RidesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			rides := h.Service.GetAllRides()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(rides)
			w.WriteHeader(http.StatusOK)
		}

	case http.MethodPost:
		{
			newRide := Ride{}
			err := json.NewDecoder(r.Body).Decode(&newRide)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			done := h.Service.CreateRide(newRide)
			if !done {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newRide)
			w.WriteHeader(http.StatusCreated)
		}
	case http.MethodPatch:
		{

		}
	case http.MethodDelete:
		{
			idStr := r.URL.Query().Get("ride_id")
			rodeID, err := uuid.Parse(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			done := h.Service.DeleteBooking(rodeID)
			if !done {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func (h *Handler) BookingsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			bookings := h.Service.GetAllBookings()
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(bookings)
			w.WriteHeader(http.StatusOK)
		}
	case http.MethodPost:
		{
			newBooking := Booking{}
			err := json.NewDecoder(r.Body).Decode(&newBooking)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			done := h.Service.CreateBooking(newBooking)
			if !done {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(newBooking)
			w.WriteHeader(http.StatusCreated)
		}
	case http.MethodPatch:
		{

		}
	case http.MethodDelete:
		{
			idStr := r.URL.Query().Get("booking_id")
			bookingID, err := uuid.Parse(idStr)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			done := h.Service.DeleteBooking(bookingID)
			if !done {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusNoContent)
		}
	}
}

func main() {
	h := NewHandler()
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/users", h.UsersHandler)
	http.HandleFunc("/rides", h.RidesHandler)
	http.HandleFunc("/bookings", h.BookingsHandler)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}
