package main

import (
	"fmt"
	"github.com/M1racle-Heen/CDEKBackend-Task/Client"
	"github.com/M1racle-Heen/CDEKBackend-Task/ResponseFolder"
	"log"
)

func main() {
	cdek, err := Client.NewClient("EMscd6r9JnFiQ3bLoyjJY6eM78JrJceI", "PjLZkKBHEiLK3YsjtNrt3TGNG0ahs3kG", "https://api.edu.cdek.ru/v2")

	if err != nil {
		fmt.Println(err)
	}
	// calculate prices
	prices, err := cdek.Calculate("Россия, г. Москва, Cлавянский бульвар д.1", "Россия, Воронежская обл., г. Воронеж, ул. Ленина д.43", ResponseFolder.Packages{
		Weight: 10,
		Length: 10,
		Width:  20,
		Height: 30,
	})
	if err != nil {
		log.Fatalf("Failed to calculate prices: %v", err)
	}
	_ = prices
	// print calculated prices
	for i, price := range prices {
		fmt.Printf("%d and %v\n", i, price)
	}
}
