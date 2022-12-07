package Pt2

import (
	Struct "Day2/Struct"
	"bufio"
	"os"
)

func Pt2(input_file_name string) (int, error) {
	// Open file and defer the close
	input_file, err := os.Open(input_file_name)
	if err != nil {
		return 0, err
	}
	defer input_file.Close()

	// output Variable
	var points int = 0

	// Read file line by line
	file_scanner := bufio.NewScanner(input_file)
	for file_scanner.Scan() {
		opponent, response := file_scanner.Text()[0], file_scanner.Text()[2]

		// Check appropriate response and add points
		switch response {
		case 'Z':
			points += Struct.Points['W']
			points += Struct.Points[Struct.Rules[opponent]]
			break
		case 'Y':
			points += Struct.Points['D']
			points += Struct.Points[Struct.Ties[opponent]]
			break
		case 'X':
			points += Struct.Points[Struct.Lose[opponent]]
			break
		}
	}
	return points, nil
}
