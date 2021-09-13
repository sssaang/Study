package main

type Reservation struct {
	audience *Audience
	numAudience int
	screening *Screening
}

func NewReservation(audience *Audience, numAudience int, screening *Screening) *Reservation {
	return &Reservation{
		audience: audience,
		numAudience: numAudience,
		screening: screening,
	}
}