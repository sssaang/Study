package main

import (
	"time"
)

type Screening struct {
	movie *Movie
	sequence int
	screenTime time.Time
}

func NewScreening(movie *Movie, sequence int, screenTime time.Time) *Screening {
	return &Screening{
		movie: movie,
		sequence: sequence,
		screenTime: screenTime,
	}
}

func (s *Screening) GetScreenTime() time.Time {
	return s.screenTime
}

func (s *Screening) GetSequence() int {
	return s.sequence
}