package main

import (
	"fmt"
	"log"
)

func main() {
	d := NewDeck()
	fmt.Println("Original Deck :")
	d.Print()
	err := d.SaveToFile("deck")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("-----------")
	hand, d := d.Deal(5)
	fmt.Println("Hand :")
	hand.Print()
	fmt.Println("-----------")
	fmt.Println("Remaining Deck :")
	d.Print()
	fmt.Println("-----------")
	fmt.Println("Shuffled Deck :")
	d, err = ShuffleDeck(d)
	if err != nil {
		log.Fatal(err)
	}
	d.Print()
	fmt.Println("-----------")
	fmt.Println("Shuffled Deck :")
	d, err = ShuffleDeck(d)
	if err != nil {
		log.Fatal(err)
	}
	d.Print()
}
