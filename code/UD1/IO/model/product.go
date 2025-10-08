package models

import("fmt")

type Product struct {
	ID       string
	Name     string
	Category string
	Price    float64
	Stock    int
}

func ParseProductFromText(text string) Product, error {

}

type InvalidArgumentsLengthError struct {
	code int,
	message string
}

func (err *InvalidArgumentsLengthError) Error() string {
	
}