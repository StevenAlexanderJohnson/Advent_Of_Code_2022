package pt2

import (
	Helpers "Day6/Helpers"
	"bufio"
	"os"
)

// This function is exactly the same as Part 1 but the window is larger.
func Pt2(file_name string) (int, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}

	defer input_file.Close()

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)

	scanner.Scan()

	line := scanner.Text()

	for i := 0; i < len(line)-14; i++ {
		window := line[i : i+14]
		if Helpers.Check_Duplicates(window) {
			continue
		}
		return i + 14, nil
	}

	return 0, nil
}
