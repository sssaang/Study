package main

import (
	"errors"
)

type Bag struct {
	cash float32
	ticket *Ticket
	invitation *Invitation
}

func newBag(cash float32, ticket *Ticket, invitation *Invitation) *Bag {
	return &Bag{
		cash: cash,
		ticket: ticket,
		invitation: invitation,
	}
}

func (b *Bag) minusCash(ticketPrice float32) error {
	if b.cash < ticketPrice {
		return errors.New("not enough cash")
	}

	b.cash -= ticketPrice
	return nil
}

func (b *Bag) checkInvitation() bool {
	if b.invitation == nil || b.invitation.validateInvitation() {
		return false
	}

	return true
}

func (b *Bag) setTicket(ticket *Ticket) {
	b.ticket = ticket
}