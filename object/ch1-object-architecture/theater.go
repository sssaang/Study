package main

import "fmt"

type Theater struct {
	ticketSeller *TicketSeller
}

func newTheater() *Theater {
	return &Theater{
		ticketSeller: newTicketSeller(),
	}
}

func (th *Theater) enter(audience *Audience) {
	if audience != nil {
		fmt.Println("audience enterring")
		err := th.ticketSeller.sellTicketToAudience(audience)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("audience enterred")
	}
}
