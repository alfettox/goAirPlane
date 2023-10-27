package database

// Define the Booking data structure
type Booking struct {
    PassengerName string
    FlightNumber  string
    DepartureCity string
    ArrivalCity   string
    SeatNumber    string
}

var Bookings = []Booking{
    {
        PassengerName: "Alessia",
        FlightNumber:  "AA415-656",
        DepartureCity: "Montréal",
        ArrivalCity:   "Roma",
        SeatNumber:    "44b",
    },
    {
        PassengerName: "Marcella",
        FlightNumber:  "CA615-656",
        DepartureCity: "Montréal",
        ArrivalCity:   "Rio de Janeiro",
        SeatNumber:    "12c",
    },
}
