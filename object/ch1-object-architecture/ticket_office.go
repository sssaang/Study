package main


type TicketOffice struct {
	tickets []*Ticket
}

func newTicketOffice() *TicketOffice {
	const NUM_TICKETS = 10
	newTickets := []*Ticket{}

	for i:=0; i<len(newTickets); i++ {
		newTickets = append(newTickets, newTicket())
	}
	return &TicketOffice{
		tickets: newTickets,
	}
}

func (to *TicketOffice) getTicket() *Ticket {
	if len(to.tickets) == 0 {
		return nil
	}
	ticket := to.tickets[0]
	to.tickets = to.tickets[1:]
	return ticket
}
