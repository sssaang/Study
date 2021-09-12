package main

import (
	"time"
)

type Invitation struct {
	inv_date time.Time
}

func newInvitation(inv_dat_str string) (*Invitation, error) {
	layout := "2006-Jan-02"
	inv_date, err := time.Parse(layout, inv_dat_str)
	if err != nil {
		return nil, err
	}
	return &Invitation{
		inv_date: inv_date,
	}, nil
}

