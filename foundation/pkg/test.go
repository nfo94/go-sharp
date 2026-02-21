package pkg

import "fmt"

// Uppercased, meaning it can be "seen" by outer packages
func Test() {
	fmt.Print("Test\n")
}
