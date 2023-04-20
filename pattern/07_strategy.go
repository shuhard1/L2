package pattern

// бизнес логика обработки заказа, которую не надо связывать с логикой оплаты заказа
func ProcessOrder(product string, payment Payment) {
	// ... implementation
	//логику оплаты заказа скрыли
	err := payment.Pay()
	if err != nil {
		return
	}
}

// ----

type Payment interface {
	Pay() error
}

// ----

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

//логика оплаты заказа
func (p *cardPayment) Pay() error {
	// ... implementation
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

//логика оплаты заказа
func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

//логика оплаты заказа
func (p *qiwiPayment) Pay() error {
	// ... implementation
	return nil
}
