package main

import (
	Pt1 "Day7/Pt1"
	"flag"
	"fmt"
)

func main() {
	var pt1 bool
	var file_name string
	flag.BoolVar(&pt1, "pt1", false, "Run part 1")
	flag.StringVar(&file_name, "file", "input.txt", "Name of the input file")
	flag.Parse()

	if pt1 {
		output, err := Pt1.Pt1(file_name)
		if err != nil {
			fmt.Printf("Error: %s", err)
			return
		}
		fmt.Printf("Output: %d", output)
	}
}
