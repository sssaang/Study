package main

import "fmt"

func main() {
	espressoWithMochaSoy := &Soy{
		beverage: &Mocha {
			beverage: &Houseblend{},
		},
	}
	fmt.Println(espressoWithMochaSoy.description(), "$", espressoWithMochaSoy.cost())
	houseBlendWithMochaMilk := &Mocha{
		beverage: &Milk{
			beverage: &Houseblend{},
		},
	}
	fmt.Println(houseBlendWithMochaMilk.description(), "$", houseBlendWithMochaMilk.cost())
}