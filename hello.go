package main

import (
	"fmt"
	"math"
	. "math/big"
)

func main() {
	b := new(Int).SetUint64(math.MaxUint64)
	b.Add(b, b)

	fmt.Println(b)
}
