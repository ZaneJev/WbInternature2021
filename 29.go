package main

import (
	"fmt"
	"math/big"
)

func main(){
	a, b := new(big.Int), new(big.Int)

	a.SetString("25000000000000000000", 10)
	b.SetString("24000000000000000000", 10)

	fmt.Printf("Sum is %v\nDiv is %v\nMod is %v\nSub is %v\nMul is %v\n",
		Sum(a, b), Div(a, b), Mod(a, b), Sub(a, b), Mul(a, b))
}

func Sum(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Add(a, b)
	return result
}

func Div(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Div(a, b)
	return result
}

func Mod(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mod(a, b)
	return result
}

func Sub(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Sub(a, b)
	return result
}

func Mul(a *big.Int, b *big.Int) *big.Int {
	result := new(big.Int)
	result.Mul(a, b)
	return result
}
