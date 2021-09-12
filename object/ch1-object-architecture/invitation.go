package main

import "time"

type Invitation struct {
	exp_date time.Time
}

func newInvitation(exp_date string) (*Invitation, error) {
	layout := "2006-Jan-02"
	inv_date, err := time.Parse(layout, exp_date)
	if err != nil {
		return nil, err
	}
	
	return &Invitation{
		exp_date: inv_date,
	}, nil
}

func (inv *Invitation) validateInvidation() bool {
	today := time.Now()
	return !today.Before(inv.exp_date) && !today.After(inv.exp_date)
}