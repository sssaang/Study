package main

import "errors"

type DiscountPolicy interface {
	CalculateDiscountAmount(screening *Screening) float32
}

type AmountDiscountPolicy struct {
	discountConditions []DiscountCondition
	amount float32
}

func NewAmountDiscountPolicy(amount float32, discountConditions []DiscountCondition) *AmountDiscountPolicy {
	adp := &AmountDiscountPolicy{
		amount: amount,
		discountConditions: discountConditions,
	}
	
	return adp
}

func (adp *AmountDiscountPolicy) CalculateDiscountAmount(screening *Screening) float32 {
	for _, dc := range adp.discountConditions {
		if dc.isSatisfied(screening) {
			return adp.amount
		}
	}

	return 0.0
}

type PercentDiscountPolicy struct {
	discountConditions []DiscountCondition
	percent float32
}

func NewPercentDiscountPolicy(percent float32, discountConditions []DiscountCondition) (*PercentDiscountPolicy, error) {
	if percent > 1.0 || percent < 0.0 {
		return nil, errors.New("invalid percent")
	}
	pdp := &PercentDiscountPolicy{
		percent: percent,
		discountConditions: discountConditions,
	}
	
	return pdp, nil
}

func (pdp *PercentDiscountPolicy) CalculateDiscountAmount(screening *Screening) float32 {
	for _, dc := range pdp.discountConditions {
		if dc.isSatisfied(screening) {
			return screening.GetMovieFee() * pdp.percent
		}
	}

	return 0.0
}

type NoDiscountPolicy struct {
}

func NewNoDiscountPolicy() (*NoDiscountPolicy) {
	return &NoDiscountPolicy{}
}

func (ndp NoDiscountPolicy) CalculateDiscountAmount(screening *Screening) float32 {
	return 0.0
}