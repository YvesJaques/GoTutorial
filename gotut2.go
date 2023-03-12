package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	// var vName string = "Yves"
	// var v1, v2 = 1.2, 3.4
	// var v3 = "Hello"
	// v1 = 4.5

	// int, float64, bool, string, rune
	// Default type 0, 0.0, false, ""

	// pl(reflect.TypeOf(25))
	// pl(reflect.TypeOf(3.14))
	// pl(reflect.TypeOf(true))
	// pl(reflect.TypeOf("Hello"))
	// pl(reflect.TypeOf('🦍'))

	// cV1 := 1.5
	// cV2 := int(cV1)
	// pl(cV2)

	// cV3 := "5000000"
	// cV4, err := strconv.Atoi(cV3)
	// pl(cV4, err, reflect.TypeOf(cV4))

	// cV5 := 50000000
	// cV6 := strconv.Itoa(cV5)
	// pl(cV6)

	// cV7 := "3.14"
	// if cV8, err := strconv.ParseFloat(cV7, 64); err ==
	// 	nil {
	// 	pl(cV8)
	// }

	// cV9 := fmt.Sprintf("%.2f", 3.14)
	// pl(cV9)

	// If Conditional
	// Conditional Operators : > < >= <= == !=
	// Logical Operators : && || !

	// iAge := 8
	// if (iAge >= 1) && (iAge <= 18) {
	// 	pl("Important Birthday")
	// } else if (iAge == 21) || (iAge == 50) {
	// 	pl("Important Birthday")
	// } else if iAge >= 65 {
	// 	pl("Important Birthday")
	// } else {
	// 	pl("Not an Important Birthday")
	// }

	// pl("!true", !true)

	// fmt.Printf("%s %d %c %f %t %o %x\n", "Stuff", 1, 'A', 3.14, true, 1, 1)
	// fmt.Printf("%9f\n", 3.14)
	// fmt.Printf("%.2f\n", 3.141592)
	// fmt.Printf("%9.f\n", 3.141592)
	// sp1 := fmt.Sprintf("%9.f\n", 3.141592)
	// pl(sp1)

	// Receive customer data (Their age)
	// What is your age
	fmt.Print("What is your age? ")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	// trim whitespace from input
	age, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("You are:", age, "years old")
	}

	if age < 5 {
		pl("Too young for school")
	} else if age == 5 {
		pl("Go to Kindergarten")
	} else if age > 5 && age <= 17 {
		fmt.Printf("Go to grade %d", age-5)
		fmt.Println()
	} else {
		pl("Go to college")
	}
}
