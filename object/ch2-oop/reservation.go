package main

type Reservation struct {
	numAudience int
	screening *Screening
}

func NewReservation(numAudience int, screening *Screening) *Reservation {
	return &Reservation{
		numAudience: numAudience,
		screening: screening,
	}
}