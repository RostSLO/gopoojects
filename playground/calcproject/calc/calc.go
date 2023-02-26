package calc

//import "fmt"

func ExecuteOperation(op Calcer) float32 {

	return op.Calc()

}

type Calcer interface {
	Calc() float32
}

type Addition [2]float32

type Substraction [2]float32

type Multiplication [2]float32

type Division [2]float32

func (ad *Addition) Calc() float32 {

	return (*ad)[0] + (*ad)[1]

}

func (ad *Substraction) Calc() float32 {

	return (*ad)[0] - (*ad)[1]

}

func (ad *Multiplication) Calc() float32 {

	return (*ad)[0] * (*ad)[1]

}

func (ad *Division) Calc() float32 {

	return (*ad)[0] / (*ad)[1]

}
