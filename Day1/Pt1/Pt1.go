package pt1

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Elf struct {
	number   int
	calories int64
}

func Pt1() {
	// Open file and defer the closing
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Initialize variables to track data
	var elfs []Elf = make([]Elf, 0)
	var index int = 0
	var calorie_total int64 = 0
	var largest_total int64 = 0

	// Read file line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Empty line means that we are done with the current elf
		if scanner.Text() == "" || scanner.Text() == "\n" {
			elfs = append(elfs, Elf{number: index, calories: calorie_total})
			index++
			if calorie_total > largest_total {
				largest_total = calorie_total
			}
			calorie_total = 0
			continue
		}

		// Count the amount of calories elf is carrying
		calories, err := strconv.ParseInt(scanner.Text(), 10, 32)
		if err != nil {
			fmt.Println(err)
			return
		}
		calorie_total += calories
	}

	elfs = append(elfs, Elf{number: index, calories: calorie_total})

	fmt.Println(largest_total)
	return
}
