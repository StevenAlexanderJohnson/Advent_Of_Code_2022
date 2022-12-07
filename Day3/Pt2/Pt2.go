package Pt2

import (
	Search "Day3/Search"
	Sort "Day3/Sort"
	"bufio"
	"os"
)

func Pt2(file_name string) (int, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}
	defer input_file.Close()

	scanner := bufio.NewScanner(input_file)

	output := 0
	elf_number := 0
	elf_items := make([][]byte, 3)
	for scanner.Scan() {
		elf_items[elf_number] = Sort.Insertion_Sort([]byte(scanner.Text()))
		if elf_number == 2 {
			// Iterate over the fist elf's items
			for _, char := range elf_items[0] {
				// See if the second elf has the same item
				_, err := Search.Binary_Search(elf_items[1], char)
				if err != nil {
					continue
				}
				// See if the third elf has the same item
				_, err = Search.Binary_Search(elf_items[2], char)
				if err != nil {
					continue
				}
				var break_value bool = false
				switch int(char) >= int('a') {
				case true:
					output += int(char - 'a' + 1)
					break_value = true
					break
				case false:
					output += int(char - 'A' + 27)
					break_value = true
					break
				default:
					break_value = true
					break
				}
				if break_value {
					break
				}
			}
		}
		elf_number = (elf_number + 1) % 3
	}

	return output, nil
}
