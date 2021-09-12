package main

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
		err := a.bag.minusCash(ticket.price)
		if err != nil {
			return err
		}	
	}
	
	a.bag.setTicket(ticket)
	return nil
}