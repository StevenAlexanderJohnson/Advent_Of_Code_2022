package Search

import (
	"testing"
)

// List must be sorted from greatest byte value to least
// This fails because it is in reverse order
func TestBinary_Search(t *testing.T) {
	input_string := []byte("abcdefgh")
	var target byte = 'c'
	_, err := Binary_Search(input_string, target)
	if err != nil {
		t.Failed()
	}
}

// Test that binary search works with duplicates
func TestBinary_Search_Duplicate(t *testing.T) {
	input_string := []byte("aaabbbccc")
	var target byte = 'b'
	actual, err := Binary_Search(input_string, target)
	if err != nil {
		t.Failed()
	}
	// We don't care which one it finds, just that it finds one
	if !(actual >= 3) && !(actual <= 5) {
		t.Errorf("Was not able to find the target byte")
	}
}

// Test that capital letters are handled differently than lower case letters
func TestCount_Instance_Captial(t *testing.T) {
	input_string := []byte("StevenJonson")
	var target byte = 'S'
	expected := 1
	actual := Count_Instances(input_string, target)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

// Test that capital letters are handled differently than lower case letters
func TestCount_Instance_Lower(t *testing.T) {
	input_string := []byte("StevenJohnson")
	var target byte = 's'
	expected := 1
	actual := Count_Instances(input_string, target)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

// Make sure it finds duplicates
func TestCount_Instance_Multiple(t *testing.T) {
	input_string := []byte("StevenJonson")
	var target byte = 'n'
	expected := 3
	actual := Count_Instances(input_string, target)
	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
