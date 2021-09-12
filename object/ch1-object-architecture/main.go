package main

func main() {
	theater := newTheater()
	audience := newAudience(newBag(10.0, nil, nil))
	theater.enter(audience)
}