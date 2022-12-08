package pt2

import (
	Structs "Day4/Structs"
	"bufio"
	"os"
	"strings"
)

func Pt2(file_name string) (int, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}
	defer input_file.Close()

	var count int = 0
	var assignments []Structs.Assignment = make([]Structs.Assignment, 2)

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)

	for scanner.Scan() {
		// Read the line and split it into two assignments
		var line string = scanner.Text()
		assignment_definitions := strings.Split(line, ",")

		// Parse the assignments into the list of assignments
		for i, assignment_definition := range assignment_definitions {
			assignments[i].From, assignments[i].To, err = Structs.Parse_Assignment(assignment_definition)
			// If there is an error parsing an assignment return an error
			if err != nil {
				return 0, err
			}
		}

		// Check if there is any overlap between the assignments
		if assignments[0].From <= assignments[1].From && assignments[0].To >= assignments[1].From {
			count += 1
		} else if assignments[1].From <= assignments[0].From && assignments[1].To >= assignments[0].From {
			count += 1
		}
	}

	return count, nil
}
