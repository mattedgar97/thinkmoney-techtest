package main

type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

type Checkout struct {
	CurrentPrice int
}

func (c Checkout) Scan(SKU string) (err error) {
	return nil
}

func (c Checkout) GetTotalPrice() (totalPrice int, err error) {
	return 0, nil
}
