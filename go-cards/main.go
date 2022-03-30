package main

import (
	"fmt"
	"log"
)

func main() {
	d := newDeck()
	fmt.Println("Original Deck :")
	d.print()
	err := d.saveToFile("deck")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----------")
	hand, d := d.deal(5)
	fmt.Println("Hand :")
	hand.print()
	fmt.Println("-----------")
	fmt.Println("Remaining Deck :")
	d.print()
}
