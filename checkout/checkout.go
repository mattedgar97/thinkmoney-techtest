package main

import "errors"

type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

type Checkout struct {
	CurrentPrice int
	Prices       map[string]int
}

var prices = map[string]int{"A": 50, "B": 30, "C": 20, "D": 15}

func (c *Checkout) init() {
	c.Prices = prices
}

// Scan takes an SKU and adds it to the checkout's held current price
func (c *Checkout) Scan(SKU string) (err error) {
	if len(c.Prices) < 1 {
		err = errors.New("checkout prices have not been set")
	}
	c.CurrentPrice = c.CurrentPrice + c.Prices[SKU]
	return err
}

func (c Checkout) GetTotalPrice() (totalPrice int, err error) {
	return c.CurrentPrice, nil
}
