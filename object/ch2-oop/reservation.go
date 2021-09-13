package main

type Reservation struct {
	audience *Audience
	numAudience int
	fee float32
	screening *Screening
}

func NewReservation(audience *Audience, numAudience int, fee float32, screening *Screening) *Reservation {
	return &Reservation{
		audience: audience,
		numAudience: numAudience,
		fee: fee,
		screening: screening,
	}
}
