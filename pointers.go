package main

// import (
// 	"fmt"
// )

// var pl = fmt.Println

// func changeVal2(myPtr *int) {
// 	*myPtr = 12
// }

// func doubleArrayPointerValues(myArray *[4]int) {
// 	for i := 0; i < len(myArray); i++ {
// 		myArray[i] *= 2
// 	}
// }

// func getAverage(nums ...float64) float64 {
// 	var sum float64 = 0
// 	for _, num := range nums {
// 		sum += num
// 	}
// 	return sum / float64(len(nums))
// }

// func main() {
// 	f4 := 10
// 	pl("f4:", f4)
// 	pl("f4 Address:", &f4)
// 	var f4Ptr *int = &f4
// 	pl("f4 Address:", f4Ptr)
// 	pl("f4 Value:", *f4Ptr)
// 	*f4Ptr = 11
// 	pl("f4 Value:", *f4Ptr)

// 	pl("f4 Before Function:", f4)
// 	changeVal2(f4Ptr)
// 	pl("f4 After Function:", f4)

// 	pArr := [4]int{1, 2, 3, 4}
// 	doubleArrayPointerValues(&pArr)
// 	pl(pArr)

// 	iFloatSlice := []float64{11, 13, 17}
// 	fmt.Printf("Average: %.3f\n", getAverage(iFloatSlice...))
// }
