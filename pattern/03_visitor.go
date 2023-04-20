package pattern

import (
	"fmt"
)

type Element interface {
	Accept(Visitor)
}

type ConcreteElementA struct{}

// Второй, не менее важный, этап – добавление метода accept в интерфейс фигуры.
func (e *ConcreteElementA) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementA.Accept()")
	visitor.VisitA(e)
}

type ConcreteElementB struct{}

// Второй, не менее важный, этап – добавление метода accept в интерфейс фигуры.
func (e *ConcreteElementB) Accept(visitor Visitor) {
	fmt.Println("ConcreteElementB.Accept()")
	visitor.VisitB(e)
}

// сперва мы определяем интерфейс посетителя следующим способом:
type Visitor interface {
	//Функции VisitA(), VisitB() позволят нам добавлять функционал для квадратов,
	//кругов и треугольников соответственно.
	VisitA(*ConcreteElementA)
	VisitB(*ConcreteElementB)
}

type ConcreteVisitor struct{}

func (v *ConcreteVisitor) VisitA(element *ConcreteElementA) {
	fmt.Println("ConcreteVisitor.VisitA()")
}

func (v *ConcreteVisitor) VisitB(element *ConcreteElementB) {
	fmt.Println("ConcreteVisitor.VisitB()")
}

func main() {
	visitor := new(ConcreteVisitor)
	elementA := new(ConcreteElementA)
	elementB := new(ConcreteElementB)
	elementA.Accept(visitor)
	elementB.Accept(visitor)
}
