package main

// name of package
// every project needs 'main' if you want to be executable

// `go run main.go` (runs)
// `go run .`
// `go build main.go` (builds exe)
// `go build` (requires 'go.mod')
// `go mod init` (makes 'go.mod')
// `go mod tidy` (updates 'go.mod'/'go.sum' with dependencies)

import (
	"errors"
	"fmt"
	"math"
	"math/rand"

	"rsc.io/quote"
)

func main() {

	fmt.Println(quote.Go())

	var explicitNum int = 42
	implicitNum := 5.3692

	fmt.Printf("Numbers %.2f %d \n", implicitNum, explicitNum)

	rn := rand.Float32()
	if rn < 0.5 {
		fmt.Println("Heads!")
	} else if rn > 0.5 {
		fmt.Println("Tails!")
	} else {
		fmt.Println("Nah.... nah that's wild!")
	}

	// arrays (size limited)
	var arr0 [3]int
	arr0[1] = 9
	fmt.Println(arr0)

	arr1 := [3]int{1, 2, 3}
	fmt.Println(arr1)

	// slices (dynamic)
	arr2 := []int{}
	fmt.Println(arr2)
	arr2 = append(arr2, 5, 3)
	fmt.Println(arr2)
	arr2 = append(arr2, arr2...)
	fmt.Println(arr2)
	fmt.Println(len(arr2))
	// remove doesn't exist,
	// here are some implementations (see below)
	arr2 = remove_unordered(arr2, 0)
	fmt.Println(arr2)
	arr2 = remove_ordered(arr2, 0)
	fmt.Println(arr2)

	// maps (dicts)
	numwords := make(map[int]string)
	numwords[0] = "zero"
	numwords[1] = "one"
	numwords[2] = "two"
	numwords[3] = "three"
	fmt.Println(numwords)
	fmt.Println(numwords[2])

	delete(numwords, 0)
	for k := range numwords {
		fmt.Print(k, " ")
	}
	fmt.Println()

	for k, v := range numwords {
		fmt.Println(k, v)
	}

	// loops (only for *)
	for i := 0; i <= 5; i++ {
		fmt.Print(i, " ")
	}
	fmt.Println()

	for i := 0.0; i < 0.9; {
		fmt.Print("wah ")
		i += rand.Float64() / 2
	}
	fmt.Println()

	loop := true
	for loop { // * while loop
		fmt.Print("ow")
		loop = rand.Float32() < 0.7
	}
	fmt.Println("o")

	// functions & errors
	x, err := sqrt(5)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(x)
	}

	x, err = sqrt(-3)
	fmt.Println(err)

	// structs (objects)
	spot := newCat("Spot")
	fmt.Println(*spot) // the * is to dereference
	// mostly not needed but spot is technically a reference
	moro := cat{name: "Moro", age: 15, happy: true}
	moro.pet()
	moro.pet()
	moro.pet()
}

// #### function w/ error
func sqrt(n float64) (float64, error) {
	if n < 0 {
		return -1, errors.New("undefined for negative numbers :<")
	} else {
		return math.Sqrt(n), nil
	}
}

// #### struct
type cat struct {
	name  string
	age   int
	happy bool
}

// struct func
// (the * prefix makes it a pointer receiver type, passed if needed with & prefix)
// this means it won't copy the struct when passed, more efficient and allows for mutation
func (c *cat) pet() {
	if c.happy {
		fmt.Println(c.name + ": prr...")
		c.happy = false
	} else {
		fmt.Println(c.name + ": meow.")
		c.happy = true
	}
}

// struct constructor
func newCat(name string) *cat { return &cat{name: name, age: 0, happy: true} }

// #### slice remove implementations
func remove_ordered(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
} // slower but retains order O(n)

func remove_unordered(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
} // fast but disordered O(1)
//----
