// Homework 2: Object Oriented Programming
// Due February 7, 2017 at 11:59pm
package main

import (
	"fmt"
	"log"
)

func main() {
	// Feel free to use the main function for testing your functions
	world := struct {
		English string
		Spanish string
		French  string
	}{
		"world",
		"mundo",
		"monde",
	}
	fmt.Printf("Hello, %s/%s/%s!", world.English, world.Spanish, world.French)
}

// Price is the cost of something in US cents.
type Price int64

// String is the string representation of a Price
// These should be represented in US Dollars
// Example: 2595 cents => $25.95
func (p Price) String() string {
	return fmt.Sprintf("$%.2f", float32(p)/100)
}

// Prices is a map from an item to its price.
var Prices = map[string]Price{
	"eggs":          219,
	"bread":         199,
	"milk":          295,
	"peanut butter": 445,
	"chocolate":     150,
}

// RegisterItem adds the new item in the prices map.
// If the item is already in the prices map, a warning should be displayed to the user,
// but the value should be overwritten.
// Bonus (1pt) - Use the "log" package to print the error to the user
func RegisterItem(prices map[string]Price, item string, price Price) {
	if p, exist := prices[item]; exist {
		log.Printf("Item has already registered. Updating price from %s to %s\n",
			p.String(), price.String())
	} else {
		fmt.Printf("New item registring. %s: %s\n", item, price.String())
	}
	prices[item] = price
}

// Cart is a struct representing a shopping cart of items.
type Cart struct {
	Items      []string
	TotalPrice Price
}

// hasMilk returns whether the shopping cart has "milk".
func (c *Cart) hasMilk() bool {
	for _, s := range c.Items {
		if s == "milk" {
			return true
		}
	}
	return false
}

// HasItem returns whether the shopping cart has the provided item name.
func (c *Cart) HasItem(item string) bool {
	for _, s := range c.Items {
		if s == item {
			return true
		}
	}
	return false
}

// AddItem adds the provided item to the cart and update the cart balance.
// If item is not found in the prices map, then do not add it and print an error.
// Bonus (1pt) - Use the "log" package to print the error to the user
func (c *Cart) AddItem(item string) {
	if p, exist := Prices[item]; exist {
		c.Items = append(c.Items, item)
		c.TotalPrice += p
	} else {
		log.Println("Item does not exist in price list.")
	}
}

// Checkout displays the final cart balance and clears the cart completely.
func (c *Cart) Checkout() {
	fmt.Println("Cart Checkout")
	for _, i := range c.Items {
		fmt.Println(i, Prices[i])
	}
	fmt.Println("Total:", c.TotalPrice)
	c.TotalPrice = 0
	c.Items = []string{}
}
