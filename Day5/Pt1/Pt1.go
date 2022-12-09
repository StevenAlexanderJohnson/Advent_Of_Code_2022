package pt1

import (
	Structs "Day5/Structs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(file_name string) (string, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return "", err
	}
	defer input_file.Close()

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)

	// Scann the first line to determine the number of stacks
	scanner.Scan()
	var line string = scanner.Text()
	// Allocate the array of stacks
	var stacks []Structs.Stack = make([]Structs.Stack, len(line))
	// Parse the first line to fill the stacks
	Structs.Parse_Stacks(line, stacks[:])

	// This scanner loop is used to generate the stacks
	for scanner.Scan() {
		line = scanner.Text()
		if line == "" || line == "\n" {
			break
		}
		Structs.Parse_Stacks(line, stacks[:])
	}

	// This scanner loop is used to move the containers
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
		// Subtract 1 from the stack index to get the index of the array
		err = Structs.Move_Container(stacks[:], number, from-1, to-1)
		if err != nil {
			return "", err
		}
	}

	for _, stack := range stacks {
		fmt.Printf("%c", stack.Top_container.Value)
	}
	fmt.Printf("\n")

	return "Hello, World!", nil
}
