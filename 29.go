package main

import (
	"fmt"
	"math/big"
)

func main(){
	a, b := new(big.Int), new(big.Int)
	sum, div, mod, sub, mul := new(big.Int), new(big.Int), new(big.Int), new(big.Int), new(big.Int)

	a.SetString("25000000000000000000", 10)
	b.SetString("24000000000000000000", 10)

	fmt.Printf("Sum is %v\nDiv is %v\nMod is %v\nSub is %v\nMul is %v\n",
		sum.Add(a, b), div.Div(a, b), mod.Mod(a, b), sub.Sub(a, b), mul.Mul(a, b))
}
