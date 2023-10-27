// main.go

package main

import (
    "encoding/json"
    "fmt"
    "github.com/gorilla/mux"
    "net/http"
    "goAirPlane/database"
)

func main() {
    r := mux.NewRouter()

    r.HandleFunc("/bookings", listBookings).Methods("GET")
    r.HandleFunc("/bookings/{flightNumber}", findBookingByFlight).Methods("GET")
    r.HandleFunc("/bookings/{flightNumber}", removeBookingHandler).Methods("DELETE")

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

    // Convert the matching bookings to JSON
    bookingsJSON, err := json.Marshal(matchingBookings)
    if err != nil {
        http.Error(w, "Error encoding bookings to JSON", http.StatusInternalServerError)
        return
    }

    // Set the content type to JSON
    w.Header().Set("Content-Type", "application/json")

    // Write the JSON response
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
