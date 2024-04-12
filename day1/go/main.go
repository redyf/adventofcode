// Santa is trying to deliver presents in a large apartment building, but he can't find the right floor - the directions he got are a little confusing. He starts on the ground floor (floor 0) and then follows the instructions one character at a time.
//
// An opening parenthesis, (, means he should go up one floor, and a closing parenthesis, ), means he should go down one floor.
//
// The apartment building is very tall, and the basement is very deep; he will never find the top or bottom floors.
//
// For example:
//
// (()) and ()() both result in floor 0.
// ((( and (()(()( both result in floor 3.
// ))((((( also results in floor 3.
// ()) and ))( both result in floor -1 (the first basement level).
// ))) and )())()) both result in floor -3.
// To what floor do the instructions take Santa?

package main

import (
	"bufio"
	"fmt"
	"os"
)

func NotQuiteLisp() {
	// Open the file
	file, err := os.Open("santa.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffered reader
	reader := bufio.NewReader(file)

	// Initialize counter
	counter := 0
	position := 0
	enteredBasement := false

	// Loop to read each character
	for {
		char, _, err := reader.ReadRune() // Ensure this line is inside the loop
		if err != nil {
			// Break the loop if there's an error or EOF
			break
		}

		// Increment position
		position++

		// Check if the character is '('
		if char == '(' {
			counter++
		} else if char == ')' {
			counter--
		}

		if counter == -1 && !enteredBasement {
			fmt.Println("Santa entered the basement at position", position)
			enteredBasement = true
		}
	}

	// Print the final count
	fmt.Println("Santa reached floor", counter)
}

func main() {
	NotQuiteLisp()
}
