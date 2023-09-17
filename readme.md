"This library emulates JavaScript-like array methods. It is not intended for high-performance computing or as a comprehensive data structure library. Instead, it is designed for simple use-cases and small data interactions. For high-performance and specialized data structure needs, Go has official libraries."

```go
package main

import (
	"fmt"
	"github.com/talpx0/tiiio"
)

func main() {
	arr := tiiio.Array{1, 2, 3, 4, 5}

	sumReducer := func(a tiiio.Any, b tiiio.Any) tiiio.Any {
		return a.(int) + b.(int)
	}

	sum := arr.Reduce(sumReducer, 0)
	fmt.Println("Sum:", sum) // Output should be: Sum: 15

	// Concatenating strings using Reduce
	strArr := tiiio.Array{"Hello", " ", "world", "!"}
	stringReducer := func(a tiiio.Any, b tiiio.Any) tiiio.Any {
		return a.(string) + b.(string)
	}

	concatenated := strArr.Reduce(stringReducer, "")
	fmt.Println("Concatenated string:", concatenated) // Output should be: Concatenated string: Hello world!
}
```
