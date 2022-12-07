package Pt1

import (
	Search "Day3/Search"
	Sort "Day3/Sort"
	"bufio"
	"fmt"
	"os"
)

func Pt1(file_name string) (int, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}

	defer input_file.Close()

	var sum int = 0

	scanner := bufio.NewScanner(input_file)

	for scanner.Scan() {
		var first_half []byte = []byte(scanner.Text()[:len(scanner.Text())/2])
		var second_half []byte = []byte(scanner.Text()[len(scanner.Text())/2:])
		first_half = Sort.Insertion_Sort(first_half)
		first_half = Sort.Remove_Duplicates(first_half)
		second_half = Sort.Insertion_Sort(second_half)
		second_half = Sort.Remove_Duplicates(second_half)
		for _, char := range first_half {
			index, err := Search.Binary_Search(second_half, char)
			if err == nil {
				switch int(second_half[index]) >= int('a') {
				case true:
					sum += int(second_half[index] - 'a' + 1)
					continue
				case false:
					sum += int(second_half[index] - 'A' + 27)
					continue
				default:
					fmt.Println("What the hell?")
					break
				}
			}
		}
	}
	return sum, nil
}
