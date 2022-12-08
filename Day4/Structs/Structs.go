package structs

import (
	"strconv"
	"strings"
)

type Assignment struct {
	From int
	To   int
}

func Parse_Assignment(assignment string) (int, int, error) {
	assignment_definition := strings.Split(assignment, "-")
	from, err := strconv.Atoi(assignment_definition[0])
	if err != nil {
		return 0, 0, err
	}
	to, err := strconv.Atoi(assignment_definition[1])
	if err != nil {
		return 0, 0, err
	}
	return from, to, nil
}
