package pt2

import (
	Structs "Day5/Structs"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func Pt2(file_name string) (string, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return "", err
	}
	defer input_file.Close()

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)

	scanner.Scan()
	var line string = scanner.Text()
	var stacks []Structs.Stack = make([]Structs.Stack, len(line))
	Structs.Parse_Stacks(line, stacks[:])

	for scanner.Scan() {
		line = scanner.Text()
		if line == "" || line == "\n" {
			break
		}
		Structs.Parse_Stacks(line, stacks[:])
	}

	for scanner.Scan() {
		line = scanner.Text()
		var instructions []string = strings.Split(line, " ")
		number, err := strconv.Atoi(instructions[0])
		if err != nil {
			return "", err
		}
		from, err := strconv.Atoi(instructions[1])
		if err != nil {
			return "", err
		}
		to, err := strconv.Atoi(instructions[2])
		if err != nil {
			return "", err
		}
		err = Structs.Move_Multiple_Containers(stacks[:], number, from-1, to-1)
		if err != nil {
			return "", err
		}
	}

	var output_string []rune = make([]rune, len(stacks))
	for i, stack := range stacks {
		output_string[i] = stack.Top_container.Value
	}

	return string(output_string), nil
}
