package booking

import "fmt"

//errors
var(
	ErrSeatAlreadyBooked = fmt.Errorf("Seat is already taken")
)


type Booking struct{
	ID string
	MovieID string
	SeatID string
	UserID string
	Status string
}


type BookingStore interface{
	Book(b Booking) error
	ListBookings(movieId string) []Booking
}