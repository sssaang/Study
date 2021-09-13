package main

type Audience struct {
	name string
}

func NewAudience(name string) *Audience {
	return &Audience{
		name: name,
	}
}