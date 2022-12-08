package pt1

import (
	Structs "Day4/Structs"
	"bufio"
	"os"
	"strings"
)

func Pt1(input_file_name string) (int, error) {
	input_file, err := os.Open(input_file_name)
	if err != nil {
		return 0, err
	}
	defer input_file.Close()

	var output int = 0
	var assignments []Structs.Assignment = make([]Structs.Assignment, 2)

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)
	for scanner.Scan() {
		var line string = scanner.Text()
		assignment_definitions := strings.Split(line, ",")
		for i, assignment_definition := range assignment_definitions {
			assignments[i].From, assignments[i].To, err = Structs.Parse_Assignment(assignment_definition)
			// If there is an error parsing an assignment return an error
			if err != nil {
				return 0, err
			}
		}

		if assignments[0].From <= assignments[1].From && assignments[0].To >= assignments[1].To {
			output += 1
		} else if assignments[1].From <= assignments[0].From && assignments[1].To >= assignments[0].To {
			output += 1
		}
	}

	return output, nil
}
