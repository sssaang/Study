package main

import (
	"time"
)

type TicketSeller struct {
	ticketOffice *TicketOffice
}

func newTicketSeller() *TicketSeller {
	return &TicketSeller{
		ticketOffice: newTicketOffice(),
	}
}

func (ts *TicketSeller) sellTicketToAudience(audience *Audience) error {
	invitation := audience.showInvitation()
	ticket := ts.ticketOffice.getTicket()
	err := audience.buyTicket(ticket, validateInvitation(invitation))
	if err != nil {
		return err
	}

	return nil
}

func validateInvitation(inv *Invitation) bool {
	today := time.Now()
	return !today.Before(inv.exp_date) && !today.After(inv.exp_date)
}