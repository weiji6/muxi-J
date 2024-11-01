```go
package main

import (
	"fmt"
)

func main() {
	var month int
	fmt.Scan(&month)
	switch {
	case month == 3 || month == 4 || month == 5:
		fmt.Println("Spring")
	case month == 6 || month == 7 || month == 8:
		fmt.Println("Summer")
	case month == 9 || month == 10 || month == 11:
		fmt.Println("Autumn")
	case month == 12 || month == 1 || month == 2:
		fmt.Println("Winter")
	default:
		fmt.Println("Season unknow")
	}
}
```