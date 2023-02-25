package main

import (
	"fmt"
	"goprojects/playground/calcproject/calc"
)

func main() {

	//ad := calc.Addition{1, 9}
	//calc.executeOperation(&ad)
	fmt.Println(calc.ExecuteOperation(2, 8))
	fmt.Println("Done")

}
