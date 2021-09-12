package main

func main() {
	theater := newTheater()
	audience1 := newAudience(newBag(10.0, nil, nil))
	theater.enter(audience1)
	inv, _ := newInvitation("2021-Sep-13")
	audience2 := newAudience(newBag(10.0, nil, inv))
	theater.enter(audience2)
}