package main

import "fmt"

type Employee struct {
	Age int
}

type Customer struct {
	Age int 
}

type User interface {
	Employee | Customer
}

func GetObject[T User]() T {
	return T{}
}

func main() {
	employee := GetObject[Employee]()
	customer := GetObject[Customer]()

	fmt.Println(employee)
	fmt.Println(customer)
}