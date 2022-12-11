package main

import (
	Pt1 "Day6/Pt1"
	Pt2 "Day6/Pt2"
	"flag"
	"fmt"
)

func main() {
	var pt1 bool
	var pt2 bool
	var file_name string
	flag.BoolVar(&pt1, "pt1", false, "Run Part 1")
	flag.BoolVar(&pt2, "pt2", false, "Run Part 2")
	flag.StringVar(&file_name, "file", "input.txt", "Name of the input file")
	flag.Parse()

	if pt1 {
		output, err := Pt1.Pt1(file_name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Pt1 Output: ", output)
	}

	if pt2 {
		output, err := Pt2.Pt2(file_name)
		if err != nil {
			panic(err)
		}
		fmt.Println("Pt2 Output: ", output)
	}
}
