package token

import "math/rand"

// maxIntValue specifies the upper limit of digits we allow to include within tokens.
const maxIntValue = 6

// Generate is responsible for providing a random token consisting of 4 digits.
func Generate() (result [4]int) {
	for i := 0; i < 4; i++ {
		result[i] = rand.Intn(maxIntValue)
	}

	return
}
