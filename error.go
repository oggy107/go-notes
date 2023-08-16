package main

import (
	"errors"
	"fmt"
)

/**
func main() {
	fmt.Println(div(10, 2))
	fmt.Println(div(10, 0))
	fmt.Println(div(20, 5))
	fmt.Println(div(6, 3))
	fmt.Println(div(6, 2))
	fmt.Println(div(100, 10))

	_, err := div(20, 0)

	fmt.Println(errors.Is(err, divByZero))
	fmt.Println(errors.Is(err, DivError{}))
}
**/

type DivByZero struct{}

func (e DivByZero) Error() string {
	return "Division by zero"
}

var divByZero = DivByZero{}

type DivError struct {
	Dividend int
	Divisor  int
}

func (e DivError) Error() string {
	return fmt.Sprintf("Cannot divide %d by %d, just for fun lol", e.Dividend, e.Divisor)
}

func div(dividend, divisor int) (quotient int, err error) {
	if divisor == 0 {
		return 0, divByZero
	}

	if dividend == 6 && divisor == 3 {
		return 0, DivError{Dividend: dividend, Divisor: divisor}
	}

	if dividend == 100 {
		return 0, errors.New("We do not allow division by 100")
	}

	return dividend / divisor, nil
}
