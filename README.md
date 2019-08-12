# proverbs

## Purpose

`proverbs` helps me explaining both modules and vanity import paths in my course [Master Go](https://appliedgo.com/p/mastergo). 

You can use it for integrating Go proverbs into your code.

## Usage

```go
package main

import (
	"fmt"

	"appliedgo.net/proverbs"
)

func main() {
	fmt.Println("Random proverb:", proverbs.Random())
	for i := 0; i < 19; i++ {
		fmt.Printf("Proverb no %d: %s\n", i+1, proverbs.No(i))
	}
}
```

