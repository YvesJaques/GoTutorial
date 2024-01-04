package main

import (
	stuff "example/project/mypackage"
	"fmt"
	"reflect"
)

var pl = fmt.Println

func main() {
	fmt.Println("Hello", stuff.Name)
	intArr := []int{2, 3, 5, 7, 11}
	strArr := stuff.IntArrToStringArr(intArr)
	fmt.Println(strArr)
	fmt.Println(reflect.TypeOf(strArr))
}
