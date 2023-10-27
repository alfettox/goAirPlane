package main

import (
	"encoding/json"
	"fmt"
	"goAirPlane/database"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/bookings", listBookings).Methods("GET")
	r.HandleFunc("/booking/{flightNumber}", findBookingByFlight).Methods("GET")
	r.HandleFunc("/bookings/{flightNumber}", removeBookingHandler).Methods("DELETE")
	r.HandleFunc("/booking/{flightNumber}", createBookingHandler).Methods("POST")

	http.Handle("/", r)

	fmt.Println("Server is listening on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println(err)
	}
}

func listBookings(w http.ResponseWriter, r *http.Request) {
	bookingsJSON, err := json.Marshal(database.Bookings)
	if err != nil {
		http.Error(w, "Error encoding bookings to JSON", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(bookingsJSON)
}

func findBookingByFlight(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightNumber := params["flightNumber"]

	matchingBookings := []database.Booking{}
	for _, booking := range database.Bookings {
		if booking.FlightNumber == flightNumber {
			matchingBookings = append(matchingBookings, booking)
		}
	}

	bookingsJSON, err := json.Marshal(matchingBookings)
	if err != nil {
		http.Error(w, "Error encoding bookings to JSON", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(bookingsJSON)
}

func removeBookingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightNumber := params["flightNumber"]

	for i, booking := range database.Bookings {
		if booking.FlightNumber == flightNumber {
			database.Bookings = append(database.Bookings[:i], database.Bookings[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Booking not found", http.StatusNotFound)
}

func createBookingHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	flightNumber := params["flightNumber"]

	// extract the booking details from the JSON in the request
	var newBooking database.Booking
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&newBooking); err != nil {
		http.Error(w, "Error decoding request body", http.StatusBadRequest)
		return
	}

	for _, booking := range database.Bookings {
		if booking.FlightNumber == flightNumber {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("This booking already exists!"))
			return
		}
	}

	database.Bookings = append(database.Bookings, newBooking)
	w.WriteHeader(http.StatusCreated)
}
