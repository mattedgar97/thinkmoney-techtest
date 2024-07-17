package main

import (
	"errors"
)

type ICheckout interface {
	Scan(SKU string) (err error)
	GetTotalPrice() (totalPrice int, err error)
}

type Checkout struct {
	CurrentPrice int
	Scanned      map[string]int
	Prices       map[string]Item
}

type Item struct {
	Price         int
	DiscountCount int
	DiscountPrice int
}

var pricingRules = map[string]Item{
	"A": {Price: 50, DiscountCount: 3, DiscountPrice: 130},
	"B": {Price: 30, DiscountCount: 2, DiscountPrice: 45},
	"C": {Price: 20},
	"D": {Price: 15},
}

func (c *Checkout) init() {
	c.Prices = pricingRules
	c.Scanned = make(map[string]int)
}

// Scan takes an SKU and adds it to the checkout's held current price
func (c *Checkout) Scan(SKU string) (err error) {
	if len(c.Prices) < 1 {
		err = errors.New("checkout prices have not been set")
	}

	c.Scanned[SKU]++
	c.CurrentPrice = calculateTotalPrice(c.Scanned, c.Prices)
	return err
}

// GetTotalPrice returns the current total price of all scanned items including discounts
func (c Checkout) GetTotalPrice() (totalPrice int) {
	return c.CurrentPrice
}

func calculateTotalPrice(scannedItems map[string]int, pricing map[string]Item) int {
	totalPrice := 0
	for SKU, count := range scannedItems {
		item := pricing[SKU]
		if item.DiscountCount > 0 && count >= item.DiscountCount {
			specialBundles := count / item.DiscountCount
			remainingItems := count % item.DiscountCount
			totalPrice += specialBundles * item.DiscountPrice
			totalPrice += remainingItems * item.Price
		} else {
			totalPrice += count * item.Price
		}
	}
	return totalPrice
}
