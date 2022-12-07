package guide

type Guide struct {
	Opponent byte
	Response byte
}

// Map for the rules of rock paper scissors
var Rules = map[byte]byte{
	'A': 'Y', // Rules to win
	'B': 'Z',
	'C': 'X',
}

var Ties = map[byte]byte{
	'A': 'X',
	'B': 'Y',
	'C': 'Z',
}

var Lose = map[byte]byte{
	'A': 'Z',
	'B': 'X',
	'C': 'Y',
}

var Points = map[byte]int{
	'W': 6,
	'D': 3,
	'L': 0,
	'X': 1,
	'Y': 2,
	'Z': 3,
}
