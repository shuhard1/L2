package pattern

import (
	"fmt"
)

type Handler interface {
	Request(flag bool)
}

// обработчик A
type ConcreteHandlerA struct {
	next Handler
}

// внутреняя логика обрабочтика A
func (h *ConcreteHandlerA) Request(flag bool) {
	fmt.Println("ConcreteHandlerA.Request()")
	//здесь запускает следующий обработчик
	//получается пошаговая цепь обработчиков
	if flag {
		h.next.Request(flag)
	}
}

// обработчик B
type ConcreteHandlerB struct {
	next Handler
}

// внутреняя логика обрабочтика B
func (h *ConcreteHandlerB) Request(flag bool) {
	fmt.Println("ConcreteHandlerB.Request()")
}

func main() {
	handlerA := &ConcreteHandlerA{new(ConcreteHandlerB)}
	handlerA.Request(true)
}
