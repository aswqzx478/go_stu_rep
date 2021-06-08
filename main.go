package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello playground")
	fmt.Printf("My weight on the surface of Mars is %v lbs", 149*0.3783)
	fmt.Printf("and I would be %v years old.\n ", 41*365/687)

	fmt.Printf("My weight on the surface of %v is %v lbs.\n", "Earth", 149.0)
}
