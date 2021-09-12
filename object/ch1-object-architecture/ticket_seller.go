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
	if inv == nil {
		return false
	}
	layout := "2006-Jan-02"
	todayString := time.Now().Format(layout)
	today, _ := time.Parse(layout, todayString)

	return !today.Before(inv.inv_date) && !today.After(inv.inv_date)
}