package pt1

import (
	Rules "Day2/Struct"
	"bufio"
	"os"
)

func Pt1(input_file_name string) (int, error) {
	// Open input file and defer the close
	input_file, err := os.Open(input_file_name)
	if err != nil {
		return 0, err
	}
	defer input_file.Close()

	// Output variable
	var points int = 0

	// Read the file line by line
	scanner := bufio.NewScanner(input_file)
	for scanner.Scan() {
		opponent, response := scanner.Text()[0], scanner.Text()[2]

		// Determine state of the game and add the points to the total
		if Rules.Rules[opponent] == response {
			points += Rules.Points[response]
			points += Rules.Points['W']
		} else if Rules.Ties[opponent] == response {
			points += Rules.Points[response]
			points += Rules.Points['D']
		} else {
			points += Rules.Points[response]
		}
	}

	return points, nil
}
