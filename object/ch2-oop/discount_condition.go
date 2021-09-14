package main

import "time"

type DiscountCondition interface{
	isSatisfied(screening *Screening) bool
}

type SequenceDiscountCondition struct {
	sequence int
}

func NewSequenceDiscountCondition(sequence int) *SequenceDiscountCondition {
	return &SequenceDiscountCondition{
		sequence: sequence,
	}
}

func (sdc *SequenceDiscountCondition) isSatisfied(screening *Screening) bool {
	return screening.GetSequence() == sdc.sequence
}

type PeriodDiscountCondition struct {
	startTime time.Time
	endTime time.Time
}

func NewPeriodDiscountCondition(startTime time.Time, endTime time.Time) *PeriodDiscountCondition {
	return &PeriodDiscountCondition{
		startTime: startTime,
		endTime: endTime,
	}
}

func (pdc *PeriodDiscountCondition) isSatisfied(screening *Screening) bool {
	screenTime := screening.GetScreenTime()
	return screenTime.After(pdc.startTime) && screenTime.Before(pdc.endTime)
}