package task1

import "fmt"

/*
A vending machine has the following denominations:
1c, 5c, 10c, 25c, 50c, and $1. Your task is to write a program that will
be used in a vending machine to return change. Assume that the vending machine
will always want to return the least number of coins or notes. Devise a function getChange(M, P)
where M is how much money was inserted into the machine and P the price of the item selected,
 that returns an array of integers representing the number of each denomination to return.

Example:
getChange(5, 0.99) // should return [1,0,0,0,0,4]
getChange(3.14, 1.99) // should return [0,1,1,0,0,1]
getChange(3, 0.01) // should return [4,0,2,1,1,2]
getChange(4, 3.14) // should return [1,0,1,1,1,0]
getChange(0.45, 0.34) // should return [1,0,1,0,0,0]
*/

func toCents(dollars float64) int {
	return int(dollars * 100)
}

func getChange(input float64, price float64) ([]int, error) {
	var result []int

	if input <= 0 {
		return result, fmt.Errorf("wrong input")
	}

	if price <= 0 {
		return result, fmt.Errorf("wrong price")
	}

	inputInCents := toCents(input)
	priceInCents := toCents(price)

	change := inputInCents - priceInCents

	if change < 0 {
		return result, fmt.Errorf("not enough money")
	}

	dollars := change / 100
	remains := change % 100
	fifties := remains / 50
	remains = remains % 50
	quarters := remains / 25
	remains = remains % 25
	tens := remains / 10
	remains = remains % 10
	fives := remains / 5
	cents := remains % 5

	// fmt.Printf("dollars: %v\n", dollars)
	// fmt.Printf("fifties: %v\n", fifties)
	// fmt.Printf("quarters: %v\n", quarters)
	// fmt.Printf("tens: %v\n", tens)
	// fmt.Printf("fives: %v\n", fives)
	// fmt.Printf("cents: %v\n", cents)

	result = append(result, cents, fives, tens, quarters, fifties, dollars)

	// fmt.Printf("result: %v\n", result)

	return result, nil
}
