package main

import "fmt"

func main() {
	d := newDeck()
	fmt.Println("Original Deck :")
	d.print()
	fmt.Println("-----------")
	hand, d := d.deal(5)
	fmt.Println("Hand :")
	hand.print()
	fmt.Println("-----------")
	fmt.Println("Remaining Deck :")
	d.print()
}
