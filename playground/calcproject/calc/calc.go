package calc

// func executeOperation(op Calcer) {
func ExecuteOperation(x int, y int) int {

	return x + y
	//fmt.Println(op.Calc())

}

type Calcer interface {
	Calc() float32
}

type Addition [2]float32

//type Substraction float32

//type Multiplication float32

//type Division float32

func (ad *Addition) Calc() float32 {

	return (*ad)[0] + (*ad)[1]

}