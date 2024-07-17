package main

import (
	"testing"
)

type ScanCase struct {
	Case             []string
	WantCurrentPrice int
}

var scanCases = []ScanCase{
	{Case: []string{"A", "A"}, WantCurrentPrice: 100},
	{Case: []string{"A", "A", "A"}, WantCurrentPrice: 130},
	{Case: []string{"B", "B"}, WantCurrentPrice: 45}, // etc
}

func TestScan(t *testing.T) {
	for _, scanCase := range scanCases {
		checkout := Checkout{CurrentPrice: 0}
		checkout.init()
		for _, c := range scanCase.Case {
			t.Logf("scanning: %s", c)
			err := checkout.Scan(c)
			if err != nil {
				t.Errorf("scan error: %v", err)
			}
		}
		if checkout.CurrentPrice != scanCase.WantCurrentPrice {
			t.Errorf("failed: got %v, wanted %v", checkout.CurrentPrice, scanCase.WantCurrentPrice)
		}
	}

}
