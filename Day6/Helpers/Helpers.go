package helpers

// Loops over the window and adds the character to the map, if it already exists then return true.
func Check_Duplicates(window string) bool {
	var check_map map[rune]int = make(map[rune]int)
	for _, char := range window {
		_, ok := check_map[char]
		if ok {
			return true
		}
		check_map[char] = 1
	}
	return false
}
