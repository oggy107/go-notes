package main

import "fmt"

type Shape interface {
	area() float64
	parameter() float64
}

type Rect struct {
	length float64
	height float64
}

func (r Rect) area() float64 {
	return r.length * r.height
}

func (r Rect) parameter() float64 {
	return 2 * (r.length + r.height)
}

func interfaceExample() {
	var rect Shape
	rect = Rect{
		length: 6.3,
		height: 4.2,
	}

	// type assertion (used to convert interface to concrete type)
	r, ok := rect.(Rect)

	// type switch can be used when there are multiple successive type assertions
	//  e here is instance of interface and v represnets underlying type
	/**
	switch v := e.(type) {
	case email:
		return v.toAddress, v.cost()
	case sms:
		return v.toPhoneNumber, v.cost()
	default:
		return "", v.cost()
	}
	**/

	fmt.Print(r, ok)
}
