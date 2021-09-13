package main

type Movie struct {
	name string
	fee float32
	discountPolicy *DiscountPolicy	
}

func NewMovie(name string, fee float32, discountPolicy *DiscountPolicy) *Movie {
	return &Movie{
		name: name, 
		fee: fee,
		discountPolicy: discountPolicy,
	}
}

func (m *Movie) GetFee(screenning *Screening) float32 {
	// return m.discountPolicy.getDiscount(m.fee, screenTime, sequence)
}