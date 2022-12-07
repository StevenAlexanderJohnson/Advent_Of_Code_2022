package main

import (
	Pt1 "Day2/Pt1"
	Pt2 "Day2/Pt2"
)

func main() {
	points, err := Pt1.Pt1("input.txt")
	if err != nil {
		panic(err)
	}
	println(points)

	points, err = Pt2.Pt2("input.txt")
	if err != nil {
		panic(err)
	}
	println(points)
	return
}
