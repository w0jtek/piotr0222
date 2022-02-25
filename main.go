package main

import (
	"fmt"
	"strconv"

	t "github.com/w0jtek/piotr0222/token"
)

// This value specifies how many time a user can retry gessing.
const allowedAttempts = 3

func main() {
	token := t.Generate()

	for i := 0; i < allowedAttempts; i++ {
		userToken := getUserInput()
		fmt.Printf("You entered token: %v\n", userToken)

		message := compare(token, userToken)
		if message == "" {
			fmt.Println("You won!")
			break
		}
		fmt.Printf("%s\n", message)
	}

	fmt.Println("Sorry, you have failed!")
}

func getUserInput() (result [4]int) {
	const length = 4
	for i := 0; i < length; i++ {
		fmt.Printf("Please enter digit %d of %d: ", i+1, length)
		// we have to get value typed by the user
		var input string
		fmt.Scanln(&input)

		// next we need to try if the entered string is a number
		intValue, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Sorry, the entered value is not a number - try again:")
			i--
			continue
		}

		// and finally - we need to check if it is a number in scope [1,6]
		if intValue < 0 || intValue > 6 {
			fmt.Printf("The number %d is out of allowed range [1,6] - try again:\n", intValue)
			i--
			continue
		}

		result[i] = intValue
	}

	return
}

func compare(a [4]int, b [4]int) string {
	var failedPositions []int
	for i := 0; i < 4; i++ {
		if a[i] == b[i] {
			continue
		}
		failedPositions = append(failedPositions, i+1)
	}

	if failedPositions == nil {
		return ""
	}

	return fmt.Sprintf("Your guess does not match on positions: %v\n", failedPositions)
}
