package main

type DiscountCondition interface{
	isSatisfied() bool
}

type SequenceDiscountCondition struct {}
func NewSequenceDiscountCondition() *SequenceDiscountCondition {
	return &SequenceDiscountCondition{}
}

type PeriodDiscountCondition struct {}
func NewPeriodDiscountCondition() *PeriodDiscountCondition {
	return &PeriodDiscountCondition{}
}