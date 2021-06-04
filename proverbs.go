package proverbs // import "appliedgo.net/proverbs"

import "math/rand"

// Random returns a random Go Proverb.
// Ensure to initialize math/rand with a random seed value, in order to get
// true random output.
func Random() string {
	return proverbs[rand.Intn(len(proverbs)-1)]
}

// No prints out proverb no. n.
func No(n int) string {
	if n >= len(proverbs) {
		return "There is only a finite amount of Go proverbs in the universe."
	}
	return proverbs[n]
}
