package main

import "fmt"

func main() {
	//Declare a 32 bit integer without initializing - this will be given the types zero value
	var myInteger int32
	fmt.Println("Value of myInteger: ", myInteger)

	// Declare and initialise 2 float64s and show their sum
	var myFloatingPointNumber float64 = 36.333
	var myOtherFloatingPointNumber float64 = 306.96969696
	fmt.Println("Sum of my floatingpoint numbers: ", myFloatingPointNumber+myOtherFloatingPointNumber)

	// Go will figure out the type
	x, y, z := 0, 1, 2
	fmt.Printf("x: %d\ty: %d\tz: %d\n", x, y, z)

	// Example of a complex number
	myComplexNumber := 5 + 24i
	fmt.Println("Value of myComplexNumber: ", myComplexNumber)

	// Example of grouping constant declaration/initialisations
	const(
		speedOfLight	  = 299792458 // unit: m/s
		pi				  = 3.14
		myFavouriteNumber = 108
	)

	// Example of grouping variable declarations/initialisations
	var (
		a int = 0		 // an integer
		b	  = 1.6 + 3i // a complex number
		c     = 2.7     // a floating-point number
	)
	fmt.Printf("a: %v\tb: %v\tc: %v\n", a, b, c)
}
