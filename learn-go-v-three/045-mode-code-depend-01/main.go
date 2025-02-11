package main

import (
	"fmt"

	"github.com/kaif225/puppy"
)

func main() {

	s1 := puppy.Bark()

	s2 := puppy.Barks()

	fmt.Println(s1)
	fmt.Println(s2)

	// can also use tested functions
	fmt.Println(puppy.Bark())
}
