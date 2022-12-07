# Day 1
## Prompt
The prompt for this day was that you are going on an expedition with many of Santa's elves. At the start of the expedition, the elves start to take inventory of their food. The input being a text file containing one number per line. The numbers are delimited by a new line to represent the end of one elf's list and the start of another's.

# Part 1
Your job for this part is to take your puzzle input and to calculate the number of calories carried by each elf. After doing so you are to return the number of calories carried by the elf caring the most.

This is simple enough. My plan was to simply convert each line to an integer and add it to the running total for a single elf. Once I encountered a new line, I would compare the current calorie count to that of the highest recorded yet. If it is greater, than replace the current max with the new max value.

You may notice that I tracked each elf in an array of structs. I simply did this in preparation for the second part (part two of each day always builds on top of part one).

Once I finished counting the number of calories the last elf was carrying and compared the current max, I printed out the answer.

# Part 2
Much like part one, you are to count the number of calories each elf is carrying, except this time you are supposed to return the top three amounts of calories. This is where the list of structs came in use. I stopped tracking the current max and implemented an insertion sort. After creating the list, I applied the sort and took the top three elements.

I later learned more about the append method. If I had remembered at the time of writing, I probably would have created an insert function that inserted the element at the correct place instead of sorting at the end.

```Go
func Insert(elfs []Elf, insert_value Elf) {
	i := 0
	// Loop over existing elfs in the list, if the calorie count of the new elf is lower
	// than the current elf append to slice after index and append slice before the index
	for _, elf := range elfs {
		if insert_value.calories < elf.calories {
			return append(elfs[:i], append([]Elf{insert_value}, elfs[i:]...)...)
		}
		i++
	}
	// If it reached here then the new elf has more calories than anyone
	return append(elfs, insert_value)
}
```