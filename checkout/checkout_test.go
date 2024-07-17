package main

import (
	"testing"
)

type ScanCase struct {
	Case             []string
	WantCurrentPrice int
}

var (
	cases = []ScanCase{
		{Case: []string{"A", "A", "A"}, WantCurrentPrice: 130},
		{Case: []string{"B", "B"}, WantCurrentPrice: 45},
	}
)

func TestScan(t *testing.T) {
	for _, scanCase := range cases {
		checkout := Checkout{CurrentPrice: 0}
		for _, c := range scanCase.Case {
			err := checkout.Scan(c)
			if err != nil {
				t.Errorf("scan error: %v", err)
			}
		}
		if checkout.CurrentPrice != scanCase.WantCurrentPrice {
			t.Errorf("got %v, wanted %v", checkout.CurrentPrice, scanCase.WantCurrentPrice)
		}
	}

}
