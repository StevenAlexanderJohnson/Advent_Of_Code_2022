package pt1

import (
	Structs "Day7/Structs"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Pt1(file_name string) (uint64, error) {
	input_file, err := os.Open(file_name)
	if err != nil {
		return 0, err
	}
	defer input_file.Close()

	// Create a base directory to store all the other directories and files
	directory := Structs.Directory{Name: "/", Files: make(map[string]Structs.File), Directories: make(map[string]Structs.Directory)}
	// Get a pointer to the base directory to manipulate
	current_directory := &directory

	var scanner *bufio.Scanner = bufio.NewScanner(input_file)
	// Flag to determine if we are reading the output of the ls command
	var reading_ls_output bool = false
	// This is to get rid of the first line
	scanner.Scan()
	// Used for debugging
	var line_number int = 2
	var dirs int = 0
	var files int = 0
	for scanner.Scan() {
		line := scanner.Text()

		// If we are reading the output of the ls command, parse the line and add it to the current directory
		if reading_ls_output && line[0] != '$' {
			split_line := strings.Split(line, " ")
			switch split_line[0] {
			case "dir":
				err := current_directory.Add_Directory(split_line[1])
				if err != nil {
					return 0, fmt.Errorf("line number: %d\ncommand: %s\nerror: %s", line_number, line, err)
				}
				dirs++
			default:
				size, err := strconv.ParseUint(split_line[0], 10, 64)
				if err != nil {
					return 0, nil
				}
				err = current_directory.Add_File(split_line[1], size)
				if err != nil {
					return 0, fmt.Errorf("line number: %d\ncommand: %s\nerror: %s", line_number, line, err)
				}
				files++
			}
		} else if line[0] == '$' { // Else if the line starts with the '$' character, we are determing what command to run
			reading_ls_output = false
			if strings.Contains(line, "cd") {
				split_command := strings.Split(line, " ")
				current_directory, err = current_directory.CD(split_command[2])
				if err != nil {
					return 0, fmt.Errorf("line number: %d\ncommand: %s\nerror: %s", line_number, line, err)
				}
			} else if strings.Contains(line, "ls") {
				reading_ls_output = true
			}
		}
		line_number++
	}
	directory_size, _ := directory.Get_Size()
	return directory_size, nil
}
