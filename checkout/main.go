package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	checkout := Checkout{CurrentPrice: 0}
	checkout.init()
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Scan an item: ")
		item, _ := reader.ReadString('\n')
		itemTrimmed := strings.Trim(item, "\n")
		_ = checkout.Scan(itemTrimmed)
		fmt.Println("Total price:", checkout.GetTotalPrice())
	}
}
