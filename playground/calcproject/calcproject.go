package main

import (
	"fmt"

	"github.com/rostslo/goprojects/playground/calcproject/calc"
)

func main() {

	ad := calc.Addition{1, 4}
	calc.ExecuteOperation(&ad)
	//fmt.Println(calc.ExecuteOperation(2, 8))
	fmt.Println("Done")

}
