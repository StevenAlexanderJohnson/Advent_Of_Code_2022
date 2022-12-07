package pt2

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

func insertion_sort(elfs []Elf) []Elf {
	for i := 1; i < len(elfs); i++ {
		j := i
		for j > 0 && elfs[j-1].calories < elfs[j].calories {
			elfs[j-1], elfs[j] = elfs[j], elfs[j-1]
			j--
		}
	}
	return elfs
}

func Pt2() {
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
			panic(err)
		}
		calorie_total += calories
	}

	elfs = append(elfs, Elf{number: index, calories: calorie_total})

	// for _, elf := range elfs {
	// 	fmt.Printf("Elf %d is carrying %d calories\n", elf.number, elf.calories)
	// }

	elfs = insertion_sort(elfs)

	fmt.Println(elfs[0].calories + elfs[1].calories + elfs[2].calories)
	return
}
