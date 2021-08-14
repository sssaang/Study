package main

func main() {
	espressoWithMochaSoy := &Soy{
		beverage: &Mocha {
			beverage: &Houseblend{},
		},
	}
	espressoWithMochaSoy.description()
	espressoWithMochaSoy.cost()
}