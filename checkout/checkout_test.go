package main

import (
	"testing"
)

type ScanCase struct {
	Case             []string
	WantCurrentPrice int
	WantError        bool
}

var scanCases = []ScanCase{
	{Case: []string{"A", "A"}, WantCurrentPrice: 100},
	{Case: []string{"A", "A", "A"}, WantCurrentPrice: 130},
	{Case: []string{"B", "B"}, WantCurrentPrice: 45},
	{Case: []string{"E", "E", "E", "A"}, WantCurrentPrice: 50},
	{Case: []string{"A", "B", "C", "D"}, WantCurrentPrice: 115},
	{Case: []string{"B", "A", "A", "A", "B", "D"}, WantCurrentPrice: 190},
	{Case: []string{"D", "B"}, WantCurrentPrice: 45},
	{Case: []string{"E", "E", "E", "A"}, WantCurrentPrice: 50},
	{Case: []string{}, WantCurrentPrice: 0},
}

func TestScan(t *testing.T) {
	for _, scanCase := range scanCases {
		checkout := Checkout{CurrentPrice: 0}
		checkout.init()
		for _, c := range scanCase.Case {
			t.Logf("scanning: %s", c)
			err := checkout.Scan(c)
			if err != nil {
				if !scanCase.WantError {
					t.Errorf("scan error: %v", err)
				}
			}
		}
		if checkout.CurrentPrice != scanCase.WantCurrentPrice {
			t.Errorf("failed: got %v, wanted %v", checkout.CurrentPrice, scanCase.WantCurrentPrice)
		}
	}

}
