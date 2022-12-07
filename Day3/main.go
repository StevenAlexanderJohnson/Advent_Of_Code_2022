package main

import (
	Pt1 "Day3/Pt1"
	Pt2 "Day3/Pt2"
	"Day3/Search"
	"flag"
	"fmt"
)

func main() {
	var pt1_flag bool
	var pt2_flag bool
	var file_name string
	flag.BoolVar(&pt1_flag, "pt1", false, "Run Pt1")
	flag.BoolVar(&pt2_flag, "pt2", false, "Run Pt2")
	flag.StringVar(&file_name, "file", "input.txt", "File to read from")
	flag.Parse()
	Search.Binary_Search([]byte("abcdefgh"), 'c')
	if pt1_flag == true {
		fmt.Println("Starting Pt1")
		pt1_output, err := Pt1.Pt1(file_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(pt1_output)
	}

	if pt2_flag == true {
		fmt.Println("Starting Pt2")
		pt2_output, err := Pt2.Pt2(file_name)
		if err != nil {
			panic(err)
		}
		fmt.Println(pt2_output)
	}

	return
}
