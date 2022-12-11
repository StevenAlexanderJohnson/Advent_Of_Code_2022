package pt1

import (
	Helpers "Day6/Helpers"
	"bufio"
	"os"
)

func Pt1(file_name string) (int, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}

	defer input_file.Close()

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)

	scanner.Scan()

	line := scanner.Text()

	// Create a window and loop over the whole string, if no duplicates are
	// found return the index of the last character in the window.
	for i := 0; i < len(line)-4; i++ {
		window := line[i : i+4]
		if Helpers.Check_Duplicates(window) {
			continue
		}
		return i + 4, nil
	}

	return 0, nil
}
