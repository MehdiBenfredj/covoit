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
				user, err := h.Service.GetUserByEmail(email)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(user)
					w.WriteHeader(http.StatusOK)
				}

				return
			} else if idStr := r.URL.Query().Get("user_id"); idStr != "" {
				userID, err := uuid.Parse(idStr)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
				}
				user, err := h.Service.GetUserById(userID)
				if err != nil {
					w.WriteHeader(http.StatusNotFound)
				} else {
					w.Header().Set("Content-Type", "application/json")
					json.NewEncoder(w).Encode(user)
					w.WriteHeader(http.StatusOK)
				}
			}
			users, err := h.Service.GetAllUsers()
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(users)
				w.WriteHeader(http.StatusOK)
			}
		}

	case http.MethodPost:
		{
			newUser := User{}
			err := json.NewDecoder(r.Body).Decode(&newUser)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			user, err := h.Service.CreateNewUser(newUser)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(user)
				w.WriteHeader(http.StatusOK)
			}
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
			err = h.Service.DeleteUser(userID)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			} else {
				w.WriteHeader(http.StatusNoContent)
			}
		}
	}
}

func (h *Handler) RidesHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		{
			rides, err := h.Service.GetAllRides()
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
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
			ride, err := h.Service.CreateRide(newRide)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(ride)
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
			err = h.Service.DeleteBooking(rodeID)
			if err != nil {
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
			bookings, err := h.Service.GetAllBookings()
			if err != nil {
				w.WriteHeader(http.StatusNotFound)
			}
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
			booking, err := h.Service.CreateBooking(newBooking)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			} else {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(booking)
				w.WriteHeader(http.StatusCreated)
			}
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
			err = h.Service.DeleteBooking(bookingID)
			if err != nil {
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
