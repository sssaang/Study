package main

type Ticket struct {
	price float32
}

func newTicket() *Ticket {
	return &Ticket {
		price: 10.0,
	}
}

func (t *Ticket) getPrice() float32 {
	return t.price
}