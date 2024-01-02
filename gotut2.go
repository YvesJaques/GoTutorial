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

	// //x exists only in for loop scope
	// for x := 1; x <= 5; x++ {
	// 	pl(x)
	// }

	// //while loop
	// fx := 0
	// for fx < 5 {
	// 	pl(fx)
	// 	fx++
	// }

	// aNums := []int{1, 2, 3}
	// for _, num := range aNums {
	// 	pl(num)
	// }

	// xVal := 1
	// for true {
	// 	if xVal == 5 {
	// 		break
	// 	}
	// 	pl(xVal)
	// 	xVal++
	// }

	// // Looping exercise - Guess a number
	// max := 50
	// seedSecs := time.Now().Unix()
	// rand.Seed(seedSecs)
	// randNum := rand.Intn(max) + 1

	// fmt.Printf("Guess a number between 0 and %d:", max)
	// for true {
	// 	reader := bufio.NewReader(os.Stdin)
	// 	guessInput, err := reader.ReadString('\n')

	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	guessInput = strings.TrimSpace(guessInput)
	// 	guess, err := strconv.Atoi(guessInput)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	if guess == randNum {
	// 		pl("You got it!")
	// 		break
	// 	} else if guess < randNum {
	// 		pl("Guess a higher number")
	// 	} else {
	// 		pl("Guess a lower number")
	// 	}
	// }

	// // Strings & Runes
	// sV1 := "A word"
	// replacer := strings.NewReplacer("A", "Another")
	// sV2 := replacer.Replace(sV1)
	// pl(sV2)

	// pl("Length:", len(sV2))
	// pl("Contains Another:", strings.Contains(sV2, "Another"))
	// pl("o index:", strings.Index(sV2, "o"))
	// pl("Replace:", strings.Replace(sV2, "o", "0", -1))
	// sV3 := "\nSome Words \n"
	// sV3 = strings.TrimSpace(sV3)
	// pl("Split:", strings.Split("a-b-c-d", "-"))
	// pl("Lower:", strings.ToLower(sV3))
	// pl("Upper:", strings.ToUpper(sV3))
	// pl("HasPrefix:", strings.HasPrefix("tacocat", "taco"))
	// pl("HasSuffix:", strings.HasSuffix("tacocat", "cat"))

	// rStr := "abcdefg"
	// pl("Rune Count:", utf8.RuneCountInString((rStr)))
	// for i, runeVal := range rStr {
	// 	fmt.Printf("%d : %#U : %c\n", i, runeVal, runeVal)
	// }

	// // Date & Time
	// now := time.Now()
	// pl(now.Year(), now.Month(), now.Day())
	// pl(now.Hour(), now.Minute(), now.Second())
	// loc, err := time.LoadLocation("America/New_York")
	// if err != nil {
	// 	pl(err)
	// }
	// fmt.Printf("Time in New York %s\n", now.In(loc))
	// loc, _ = time.LoadLocation("Asia/Shanghai")
	// fmt.Printf("Time in Shanghai %s\n", now.In(loc))

	// locEST, _ := time.LoadLocation("EST")
	// locUTC, _ := time.LoadLocation("UTC")
	// locMST, _ := time.LoadLocation("MST")
	// fmt.Printf("EST: %s\n", now.In(locEST))
	// fmt.Printf("UTC: %s\n", now.In(locUTC))
	// fmt.Printf("MST: %s\n", now.In(locMST))

	// birthDate := time.Date(1974, time.December, 21, 11, 30, 10, 0, time.Local)
	// diff := now.Sub(birthDate)
	// fmt.Printf("Days Alive: %d days\n", int(diff.Hours()/24))
	// fmt.Printf("Hours Alive: %d hours\n", int(diff.Hours()))

	// // Arrays
	// var arr1 [5]int
	// arr1[0] = 1
	// arr2 := [5]int{1, 2, 3, 4, 5}
	// pl("Index 0:", arr2[0])
	// pl("Array Length:", len(arr2))
	// for i := 0; i < len(arr2); i++ {
	// 	pl(arr2[i])
	// }

	// for i, v := range arr2 {
	// 	fmt.Printf("Index: %d, Value: %d\n", i, v)
	// }

	// arr3 := [2][2]int{{1, 2}, {3, 4}}
	// for i := 0; i < len(arr3); i++ {
	// 	for j := 0; j < len(arr3[i]); j++ {
	// 		pl(arr3[i][j])
	// 	}
	// }

	// aStr1 := "abcde"
	// rArr := []rune(aStr1)
	// for _, v := range rArr {
	// 	fmt.Printf("Rune Array %d\n", v)
	// }
	// byteArr := []byte{'a', 'b', 'c'}
	// bStr := string(byteArr[:])
	// pl("I'm a string:", bStr)
}
