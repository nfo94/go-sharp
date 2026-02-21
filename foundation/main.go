/*
T17
This file contains the code for the "foundations" study.
*/
package main

import (
	"fmt"
	"foundation/pkg"
)

// The usual product of a Go project a single executable. You can define the platform when you're
// compiling it (x86, ARM, etc).

// All files inside the same folder need to have the same package.
// `package main` is special as it creates an executable. Other packages are library packages for
// reusable code.

// Global scope
const a = "Hello world"

// Global scope
var (
	b bool
	c int
	d string
	e float32
)

// User-defined type
type User struct {
	name string
	age  int
}

// User-defined type
type ID int

type Address struct {
	city string
}

type FullUser struct {
	name    string
	age     int
	address Address
	// Doing just Address the address means we are composing. Above, we're saying that the type
	// of the address attribute is Address
}

// An interface is a group of behaviors.
// A struct can implement a behavior from an interface.
type SomeBehaviors interface {
	Act(msg string) string
}

// Empty interface. It will suggest changing it to `any`
type ExampleInterface interface{}

func main() {
	println(a)
	// Initial default value `false`
	println(b)
	b = true
	println(b)
	println(c)
	println(d)
	println(e)

	// Local scope with shorthand of declare and assign
	var f int = 10
	println(f)

	// Create user with specific struct
	var userOne User = User{
		name: "Jane",
		age:  33,
	}
	// Can't print the struct itself
	println(userOne.name)
	println(userOne.age)

	// S
	var idOne ID = 1
	// Prints EOF
	fmt.Println(idOne)

	// The \n adds a new line and prevents the EOF to appear ("%")
	fmt.Printf("User name: %s\n", userOne.name)

	// Composite types
	// An array in Go has a pre-defined length
	var arrOne [3]int
	arrOne[0] = 1
	arrOne[1] = 1
	arrOne[2] = 1

	// Println adds a new line in the end, just like our \n before
	fmt.Println(arrOne)
	// To get the length of the array
	fmt.Println(len(arrOne))

	var arrTwo [3]int
	arrTwo[0] = 1
	arrTwo[1] = 2
	arrTwo[2] = 3
	fmt.Println(arrTwo)

	// A slice is actually a pointer to an array
	s := []int{1, 2, 3, 4, 5}
	fmt.Printf("len=%d | cap=%d | type=%T\n", len(s), cap(s), s)
	fmt.Println(s[1:5])
	fmt.Println(s[2:5])
	fmt.Println(s[:5])
	fmt.Println(s[1:4])
	fmt.Println(s[:3])
	s = append(s, 6)
	fmt.Println(s)
	// An append within the capacity modified the slice (which is an array under the hood)
	// If it exceeds the capacity Go allocates a new array (2x growth)
	fmt.Printf("len=%d | cap=%d | type=%T\n", len(s), cap(s), s)

	// Maps (dictionaries, or hash tables). map[key]value
	mapOne := map[string]int{"first": 1, "second": 2}
	fmt.Println(mapOne["first"])
	mapOne["third"] = 3
	fmt.Println(mapOne["third"])

	// Using the built in make() function
	mapTwo := make(map[string]int)
	// Empty
	fmt.Println(mapTwo)
	for long, number := range mapOne {
		fmt.Println(long, number)
	}

	// Functions
	fmt.Printf("Using the sum() function: 1 + 2 = %d\n", sum(1, 2))

	msg := "the sum of this numbers"
	// Ignoring the first returning value
	_, sumTwo := sumVariadic(msg, 1, 2, 3, 4, 5)
	// Variadic function is one that can be called with a varying number of arguments
	fmt.Printf("Using the sumVariadic() function: %s 1 + 2 + 3 + 4 + 5 = %d\n", msg, sumTwo)

	// Anonymous function. Immediate call
	func() {
		fmt.Println(sum(1, 2))
	}()

	// Obs.: closure = function + captured environment (e.g. using an external variable)

	// Going back to the user defined struct
	userTwo := User{
		name: "John",
		age:  33,
	}
	fmt.Println(userTwo.age)
	fmt.Println(userTwo.name)

	a := 10
	// Getting the value
	println(a)
	// Getting the memory address
	println(&a)

	// We're assigning the memory address of a to a variable called `pointer`, that has to be
	// of type `*int` (* is to indicate that is a pointer, and in this case a pointer to an integer)
	var pointer *int = &a
	println(pointer)

	// Reassigning the int value 20 in the pointer of the `pointer` variable, meaning we're
	// updating the value inside the pointer. We cannot use &pointer because it is an address,
	// expression. To get "inside" the address we need to use pointer.
	*pointer = 20
	println(a)

	b := &a
	println(b)  // we assigned the memory
	println(*b) // contents at that address
	println(&b) // memory address again

	// Will suggest using `any`
	var testVar interface{} = "John Doe"
	println(testVar)
	println(testVar.(string)) // type assertion
	res, ok := testVar.(int)  // Go says: "nah I can't convert this"
	println(res, ok)

	// Generics
	m3 := map[string]MyGenericNumber{"John": 10, "Jane": 20, "Smith": 30}
	// Remove the tilde from int in `GenericNumber` to see the error
	println(genericSum2(m3))

	// Importing. You need to `go mod init` since Test is in another package. If you don't do this
	// Go will search in the usr/local/go/src and won't find the package
	pkg.Test()

	// For loop
	sampleArr := [3]int{1, 2, 3}
	for _, v := range sampleArr {
		fmt.Println(v)
	}

	// If
	var defaultBool bool
	if defaultBool {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	// Switch
	switch a {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	}
}

// (u *User) is a method receiver. This is a method of the User struct
func (u *User) ChangeName(name string) string {
	// If we try to change the name it won't since we're working in a copy
	u.name = name
	return u.name
}

// Obs.: to work with the actual value we need to pass the pointer, not the name of the variable
// itself, since it will send a copy

// Regular function
func sum(a int, b int) int {
	return a + b
}

// Variadic function
func sumVariadic(msg string, numbers ...int) (string, int) {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return msg, total
}

// Generic function to work with two types, either int or float32
func genericSum1[T int | float32](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}

type MyGenericNumber int

type GenericNumber interface {
	~int | ~float32
}

// Generic function that uses the tilde to accept any type that the underlying type is also int or
// float32. Here things start to get shady IMHO
func genericSum2[T GenericNumber](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}
	return sum
}
