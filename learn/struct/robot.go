package main

import "fmt"

type Robot struct {
	ON    bool
	Ammo  int
	Power int
}

func main() {
	r := Robot{true, 150, 100}
	for i := 0; i < 5; i++ {
		r.Shoot()
		r.Ride()
	}
	fmt.Println(r.Ammo, r.Power)
}

func (r Robot) IsActive() bool {
	return r.ON
}

func (r *Robot) Shoot() bool {
	if r.IsActive() {
		r.Ammo--
		return true
	}
	return false
}

func (r *Robot) Ride() bool {
	if r.IsActive() {
		r.Power--
		return true
	}
	return false
}
