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

func (m *Movie) GetFee() float32 {
	return m.fee
}