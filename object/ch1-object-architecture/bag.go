package main

import (
	"errors"
	"time"
)

type Bag struct {
	cash float32
	ticket *Ticket
	invitation *Invitation
}

func (b *Bag) minusCash(ticketPrice float32) error {
	if b.cash < ticketPrice {
		return errors.New("not enough cash")
	}

	b.cash -= ticketPrice
	return nil
}

func (b *Bag) checkInvitation() bool {
	if b.ticket == nil {
		return false
	}

	today := time.Now()
	if today.Before(b.invitation.exp_date) || today.After(b.invitation.exp_date) {
		return false
	}

	return true
}

func (b *Bag) setTicket(ticket *Ticket) {
	b.ticket = ticket
}