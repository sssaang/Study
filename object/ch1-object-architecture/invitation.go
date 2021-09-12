package main

import "time"

type Invitation struct {
	exp_date time.Time
}

func newInvitation(exp_date string) (*Invitation, error) {
	layout := "0000-Jan-01"
	inv_date, err := time.Parse(layout, exp_date)
	if err != nil {
		return nil, err
	}
	
	return &Invitation{
		exp_date: inv_date,
	}, nil
}