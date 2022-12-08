package main

import (
	Pt1 "Day4/Pt1"
	Pt2 "Day4/Pt2"
	"flag"
	"fmt"
)

func main() {
	var pt1 bool
	var pt2 bool
	var file_name string
	flag.BoolVar(&pt1, "pt1", false, "Run Pt1")
	flag.BoolVar(&pt2, "pt2", false, "Run Pt2")
	flag.StringVar(&file_name, "file", "input.txt", "Input file name")
	flag.Parse()

	if pt1 {
		output, err := Pt1.Pt1(file_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	}

	if pt2 {
		output, err := Pt2.Pt2(file_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	}
}
