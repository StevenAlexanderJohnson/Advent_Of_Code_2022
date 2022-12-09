package main

import (
	Pt1 "Day5/Pt1"
	"flag"
)

func main() {
	var pt1 bool
	var input_file string
	flag.BoolVar(&pt1, "pt1", false, "Run part 1")
	flag.StringVar(&input_file, "file", "minified_input.txt", "Input file")
	flag.Parse()

	if pt1 {
		Pt1.Pt1(input_file)
	}

}
