package main

// import "fmt"

// var pl = fmt.Println

// type customer struct {
// 	name    string
// 	address string
// 	balance float64
// }

// func getCustomerInfo(c customer) {
// 	fmt.Printf("%s owes us %.2f\n", c.name, c.balance)
// }

// func newCustomerAdd(c *customer, address string) {
// 	c.address = address
// }

// type rectangle struct {
// 	length, height float64
// }

// func (r rectangle) Area() float64 {
// 	return r.length * r.height
// }

// type contact struct {
// 	fName string
// 	lName string
// 	phone string
// }

// type business struct {
// 	name    string
// 	address string
// 	contact
// }

// func (b business) info() {
// 	fmt.Printf("Contact at %s is %s %s\n", b.name, b.contact.fName, b.contact.phone)
// }

// func main() {
// 	var tS customer
// 	tS.name = "Tom Smith"
// 	tS.address = "5 Main St"
// 	tS.balance = 234.56

// 	getCustomerInfo(tS)
// 	newCustomerAdd(&tS, "123 South St")
// 	pl("Address:", tS.address)
// 	sS := customer{"Sally Smith", "123 Main St", 0}
// 	pl("Name:", sS.name)

// 	rect1 := rectangle{length: 10.0, height: 15.0}
// 	pl("Rectangle Area:", rect1.Area())

// 	con1 := contact{
// 		"James", "Wang", "555-1212",
// 	}

// 	bus1 := business{
// 		"ABC Corp", "123 Main St", con1,
// 	}
// 	bus1.info()
// }
