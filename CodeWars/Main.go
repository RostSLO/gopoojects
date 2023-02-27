package main

import (
	"fmt"
	"strings"
)

func main() {

	//slc := "Welcome"
	//slc := "to"
	//slc := "CodeWars"
	slc := "Hey fellow warriors"
	res := SpinWords(slc)

	fmt.Println(res)

}

func SpinWords(str string) string {

	strSlc := strings.Split(str, " ")
	res := ""
	var revertWord []byte
	for _, x := range strSlc {
		if len(x) >= 5 {
			for i := len(x) - 1; i >= 0; i-- {
				revertWord = append(revertWord, x[i])
			}
			res += string(revertWord) + " "
			revertWord = nil
		} else {
			res += x + " "
		}
	}
	return res[:len(res)-1]
} // SpinWords

/*
func FindOdd(seq []int) int {
	dict := make(map[int]int)
	for _, num := range seq {
		dict[num] = dict[num] + 1
	}
	fmt.Println(dict)
	for k, v := range dict {
		if v%2 != 0 {
			return k
		}
	}
	return 0
}


func Multiple3And5(number int) int {
	var sum int = 0

	if number < 3 {
		return 0
	}

	for i := 3; i < number; i++ {
		if i%3 == 0 {
			sum += i
		} else if i%5 == 0 {
			sum += i
		}
	}

	return sum

}
*/
