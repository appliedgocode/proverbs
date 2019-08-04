package proverbs // import "appliedgo.net/proverbs"

import "math/rand"

func Random() string {
	return proverbs[rand.Intn(len(proverbs)-1)]
}

func No(n int) string {
	if n >= len(proverbs) {
		return "There is only a finite amount of Go proverbs in the universe."
	}
	return proverbs[n]
}
