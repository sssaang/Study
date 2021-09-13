package main

import "errors"

type DiscountPolicy struct {
	discountConditions []DiscountCondition
}

type AmountDiscountPolicy struct {
	DiscountPolicy
	amount float32
}

func NewAmountDiscountPolicy(amount float32, discountConditions []DiscountCondition) *AmountDiscountPolicy {
	adp := &AmountDiscountPolicy{
		amount: amount,
	}
	adp.discountConditions = discountConditions
	return adp
}

type PercentDiscountPolicy struct {
	DiscountPolicy
	percent float32
}

func NewPercentDiscountPolicy(percent float32, discountConditions []DiscountCondition) (*PercentDiscountPolicy, error) {
	if percent > 1.0 || percent < 0.0 {
		return nil, errors.New("invalid percent")
	}
	pdp := &PercentDiscountPolicy{
		percent: percent,
	}
	pdp.discountConditions = discountConditions
	return pdp, nil
}