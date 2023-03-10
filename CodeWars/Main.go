package main

import (
	"fmt"
)

func main() {

	testData := 11
	//testData := 1675030927
	//testData := 195
	//testData := 16
	res := IsPrime(testData)

	fmt.Println(res)

}

func IsPrime(n int) bool {

	if n <= 1 {
		return false
	}

	if n == 2 || n == 3 || n == 5 || n == 7 {
		return true
	}

	for j := 2; j*j <= n; j++ {
		if n%j == 0 {
			return false
		}
	}

	return true
}

/*
func CreatePhoneNumber(numbers [10]uint) string {

	var res string = strings.Trim(strings.Join(strings.Fields(fmt.Sprint(numbers)), ""), "[]")

	return "(" + res[:3] + ") " + res[3:6] + "-" + res[6:]
}


func DigitalRoot(n int) int {

	if n <= 9 {
		return n
	}
	weAreDone := true
	res := 0

	for weAreDone {
		str := fmt.Sprintf("%v", n)
		for i := 0; i < len(str); i++ {
			num, err := strconv.Atoi(string(str[i]))
			if err == nil {
				res = res + num
			}
		}
		if res <= 9 {
			return res
		} else {
			n = res
			res = 0
		}
	}
	return 0
}

/func SpinWords(str string) string {

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
