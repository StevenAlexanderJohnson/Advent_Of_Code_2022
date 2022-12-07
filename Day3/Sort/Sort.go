package Sort

func Insertion_Sort(input_string []byte) []byte {
	for i := 1; i < len(input_string); i++ {
		for j := i; j > 0; j-- {
			if input_string[j] > input_string[j-1] {
				input_string[j], input_string[j-1] = input_string[j-1], input_string[j]
			}
		}
	}
	return input_string
}

// Removes duplicates from a sorted slice of bytes
func Remove_Duplicates(input_string []byte) []byte {
	for i := 0; i < len(input_string)-1; i++ {
		if input_string[i] == input_string[i+1] {
			input_string = append(input_string[:i], input_string[i+1:]...)
			i--
		}
	}
	return input_string
}
