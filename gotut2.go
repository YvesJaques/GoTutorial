package main

import (
	"fmt"
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

	// // Receive customer data (Their age)
	// // What is your age
	// fmt.Print("What is your age? ")
	// reader := bufio.NewReader(os.Stdin)
	// input, err := reader.ReadString('\n')
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // trim whitespace from input
	// age, err := strconv.Atoi(strings.TrimSpace(input))
	// if err != nil {
	// 	log.Fatal(err)
	// } else {
	// 	fmt.Println("You are:", age, "years old")
	// }

	// if age < 5 {
	// 	pl("Too young for school")
	// } else if age == 5 {
	// 	pl("Go to Kindergarten")
	// } else if age > 5 && age <= 17 {
	// 	fmt.Printf("Go to grade %d", age-5)
	// 	fmt.Println()
	// } else {
	// 	pl("Go to college")
	// }

	// // Math
	// pl("5 + 4 =", 5+4)
	// pl("5 - 4 =", 5-4)
	// pl("5 * 4 =", 5*4)
	// pl("5 / 4 =", 5/4)
	// pl("5 % 4 =", 5%4)
	// mInt := 1
	// mInt += 1
	// mInt = mInt + 1

	// pl("Float Precision =", 0.1111111111111111+0.1111111111111111)

	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(50) + 1
	// pl("Random:=", randNum)

	// pl("Abs(-10) = ", math.Abs(-10))
	// pl("Pow(4, 2) = ", math.Pow(4, 2))
	// pl("Sqrt(16) = ", math.Sqrt(16))
	// pl("Cbrt(8) = ", math.Cbrt(8))
	// pl("Ceil(4.4) = ", math.Ceil(4.4))
	// pl("Floor(4.4) = ", math.Floor(4.4))
	// pl("Round(4.4) = ", math.Round(4.4))
	// pl("Log2(8) = ", math.Log2(8))
	// pl("Log10(100) = ", math.Log10(100))
	// pl("Log(7.389) = ", math.Log(7.389))
	// pl("Max(5, 4) = ", math.Max(5, 4))
	// pl("Min(5,4) = ", math.Min(5, 4))

	// r90 := 90 * math.Pi / 180
	// d90 := r90 * (180 / math.Pi)
	// fmt.Printf("%f radians = %f degrees\n", r90, d90)
	// pl("Sin(90) =", math.Sin(r90))

	// fmt.Print("Enter number 1:")
	// reader := bufio.NewReader(os.Stdin)
	// number1Input, err := reader.ReadString('\n')

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// number1Input = strings.TrimSpace(number1Input)
	// number1, err := strconv.Atoi(number1Input)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Print("Enter number 2:")
	// reader = bufio.NewReader(os.Stdin)
	// number2Input, err := reader.ReadString('\n')

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// number2Input = strings.TrimSpace(number2Input)
	// number2, err := strconv.Atoi(number2Input)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Printf("%d + %d = %d\n", number1, number2, number1+number2)
}
