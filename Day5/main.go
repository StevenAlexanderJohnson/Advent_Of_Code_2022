package main

import (
	Pt1 "Day5/Pt1"
	Pt2 "Day5/Pt2"
	"flag"
	"fmt"
)

func main() {
	var pt1 bool
	var pt2 bool
	var input_file string
	flag.BoolVar(&pt1, "pt1", false, "Run part 1")
	flag.BoolVar(&pt2, "pt2", false, "Run part 2")
	flag.StringVar(&input_file, "file", "minified_input.txt", "Input file")
	flag.Parse()

	if pt1 {
		output, err := Pt1.Pt1(input_file)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	}

	if pt2 {
		output, err := Pt2.Pt2(input_file)
		if err != nil {
			panic(err)
		}
		fmt.Println(output)
	}

}
