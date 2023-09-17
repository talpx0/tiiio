"This library emulates JavaScript-like array methods. It is not intended for high-performance computing or as a comprehensive data structure library. Instead, it is designed for simple use-cases and small data interactions. For high-performance and specialized data structure needs, Go has official libraries."



```go
package main

import (
	"fmt"
	"github.com/talpx0/tiiio" // Replace with your actual package import
)

func main() {
	arr := tiiio.Array{1, 2, 3, 4, 5}

	// Sum reducer function
	sumReducer := func(a, b int) int {
		return a + b
	}

	// Using Reduce to calculate the sum of array elements
	sum := arr.Reduce(sumReducer, 0)
	fmt.Println("Sum:", sum) // Output should be: Sum: 15

	// String array
	strArr := tiiio.Array{"Hello", " ", "world", "!"}

	// String concatenation reducer function
	stringReducer := func(a, b string) string {
		return a + b
	}

	// Using Reduce to concatenate the strings
	concatenated := strArr.Reduce(stringReducer, "")
	fmt.Println("Concatenated string:", concatenated) // Output should be: Concatenated string: Hello world!
}
```
