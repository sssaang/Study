package main

import (
	"fmt"
	"time"
)

func main() {
	// discount conditions
	sdc1 := NewSequenceDiscountCondition(3)
	sdc2 := NewSequenceDiscountCondition(6)
	sdc3 := NewSequenceDiscountCondition(12)
	sdc4 := NewSequenceDiscountCondition(20)

	layout := "2006-01-02 15:04:05"
	
	nineOClock, _ := time.Parse(layout, "2021-09-13 09:00:00")
	tenOClock, _ := time.Parse(layout, "2021-09-13 10:00:00")
	elevenOClock, _ := time.Parse(layout, "2021-09-13 11:00:00")
	twelveOClock, _ := time.Parse(layout, "2021-09-13 12:00:00")
	oneOClock, _ := time.Parse(layout, "2021-09-13 13:00:00")
	twoOClock, _ := time.Parse(layout, "2021-09-13 14:00:00")

	pdc1 := NewPeriodDiscountCondition(nineOClock, twelveOClock)
	pdc2 := NewPeriodDiscountCondition(twelveOClock, twoOClock)
	

	// discount policies
	ndp := NewNoDiscountPolicy()
	adp1 := NewAmountDiscountPolicy(5.0, []DiscountCondition{sdc1, pdc1, pdc2})
	adp2 := NewAmountDiscountPolicy(3.0, []DiscountCondition{sdc2, sdc3, sdc4, pdc1})
	
	pdp1, _ := NewPercentDiscountPolicy(0.15, []DiscountCondition{sdc1, pdc1})
	pdp2, _ := NewPercentDiscountPolicy(0.18, []DiscountCondition{sdc3, pdc1})

	// movies
	titanic := NewMovie("Titanic", 15.32, adp1)
	avatar := NewMovie("Avatar", 15.99, adp2) 
	goneWithTheWind := NewMovie("Gone With The Wind", 12.99, pdp1)
  inception := NewMovie("Inception", 14.56, pdp2)
	expensiveMovie := NewMovie("Expensive Movie", 20.99, ndp)

	// screenings
	goneWithTheWind9am := NewScreening(goneWithTheWind, 1, nineOClock)
	titanic10am := NewScreening(titanic, 2, tenOClock)
	_1 := NewScreening(expensiveMovie, 3, elevenOClock)
	avatar12pm := NewScreening(avatar, 4, twelveOClock)
	_2 := NewScreening(inception, 5, oneOClock)

	fmt.Println(_1)
	fmt.Println(_2)

	// audiences
	aud1 := NewAudience("Bob Ross")
	aud2 := NewAudience("John Doe")
	aud3 := NewAudience("Jahe Doe")

	// reservations
	res1 := goneWithTheWind9am.MakeReservation(3, aud1)
	res2 := titanic10am.MakeReservation(1, aud2)
	res3 := avatar12pm.MakeReservation(2, aud3)

	fmt.Println(res1.fee)
	fmt.Println(res2.fee)
	fmt.Println(res3.fee)
}