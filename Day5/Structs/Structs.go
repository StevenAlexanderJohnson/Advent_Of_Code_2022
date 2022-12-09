package structs

import "fmt"

type Container struct {
	Value rune
	Next  *Container
}

type Empty_Container interface {
	Error() string
}

func (m *Container) Error() string {
	return "Empty Container"
}

type Stack struct {
	Top_container *Container
}

// Push a new container onto the bottom of the stack
// Used when first generating the stacks, built top to bottom
func (m *Stack) Push(value rune) {
	new_container := Container{Value: value, Next: nil}
	if m.Top_container == nil {
		m.Top_container = &new_container
		return
	}

	temp_container := m.Top_container
	for temp_container.Next != nil {
		temp_container = temp_container.Next
	}
	temp_container.Next = &new_container
}

// The line is a string of characters, each character is a container
// spaces represent absence of a container
func Parse_Stacks(line string, stacks []Stack) {
	for i, char := range line {
		if char == ' ' {
			continue
		}
		stacks[i].Push(char)
	}
}

// Pop the top container off the stack
func (m *Stack) Pop() (*Container, error) {
	temp_container := m.Top_container
	if temp_container == nil {
		return nil, Empty_Container(temp_container)
	}
	m.Top_container = temp_container.Next
	return temp_container, nil
}

// Push to top of the stack
func (m *Stack) Add(container *Container) error {
	if container == nil {
		return Empty_Container(container)
	}
	container.Next = m.Top_container
	m.Top_container = container
	return nil
}

// Move the top number containers from the from stack to the to stack
func Move_Container(stacks []Stack, number int, from int, to int) error {
	// Loop to remove correct number of containers
	for i := 0; i < number; i++ {
		temp_container, err := stacks[from].Pop()
		if err != nil {
			return fmt.Errorf("Removing Container: %s", err)
		}
		// This check is not needed, if the pop fails this code is unreachable
		// Keeping it to be safe
		err = stacks[to].Add(temp_container)
		if err != nil {
			return fmt.Errorf("Adding Container: %s", err)
		}
	}
	return nil
}
