package proverbs_test

import (
	"fmt"

	"appliedgo.net/proverbs"
)

func ExampleRandom() {
	fmt.Println(proverbs.Random())
}

func ExampleNo() {
	fmt.Println(proverbs.No(18))

	// Output:
	// Don't panic.
}
