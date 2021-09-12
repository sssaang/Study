package main

import "fmt"

type Audience struct {
	bag *Bag
}

func newAudience(bag *Bag) *Audience {
	return &Audience{
		bag: bag,
	}
}

func (a *Audience) showInvitation() *Invitation {
	return a.bag.invitation
}

func (a *Audience) buyTicket(ticket *Ticket, isInvited bool) error {
	if !isInvited {
		fmt.Println("the audience is not invited")
		err := a.bag.minusCash(ticket.getPrice())
		if err != nil {
			return err
		}
	}
	
	a.bag.setTicket(ticket)
	return nil
}