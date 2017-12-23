package main

import (
	"reflect"
	"testing"
)

func TestPriceString(t *testing.T) {
	sample := []Price{1, 2, 3, 1234, 1000}
	res := []string{"$0.01", "$0.02", "$0.03", "$12.34", "$10.00"}
	for i, s := range sample {
		if s.String() != res[i] {
			t.Errorf(`Price function String() != %v`, res[i])
		}
	}
}

func TestRegisterItem(t *testing.T) {
	sample_param1 := []string{"apple", "eggs", "lemon", "banana", "bread"}
	sample_param2 := []Price{199, 319, 299, 59, 219}
	res := map[string]Price{
		"bread":         219,
		"milk":          295,
		"peanut butter": 445,
		"chocolate":     150,
		"apple":         199,
		"lemon":         299,
		"banana":        59,
		"eggs":          319,
	}
	for i := 0; i < 5; i++ {
		RegisterItem(Prices, sample_param1[i], sample_param2[i])
	}
	equal := reflect.DeepEqual(Prices, res)
	if !equal {
		t.Errorf("output: %v\nexpecting: %v", Prices, res)
	}
}

func TestAddItem(t *testing.T) {
	var cart Cart
	sample := []string{"eggs", "apple", "milk", "green pepper"}
	res := []string{"eggs", "apple", "milk"}
	for _, s := range sample {
		cart.AddItem(s)
	}
	if !reflect.DeepEqual(cart.Items, res) {
		t.Errorf("Cart AddItem() Error")
	}
}

func TestHasMilk(t *testing.T) {
	var cart Cart
	cart.AddItem("milk")
	if !cart.hasMilk() {
		t.Errorf("Cart function hasMilk() Error.", cart)
	}
}

func TestHasItem(t *testing.T) {
	var cart Cart
	sample_add := []string{"eggs", "apple", "milk", "green pepper"}
	sample_check := []string{"apple", "bread", "peanut butter"}
	res := []bool{true, false, false}
	for _, s := range sample_add {
		cart.AddItem(s)
	}
	for i, s := range sample_check {
		if cart.HasItem(s) != res[i] {
			t.Errorf("Cart function HasItem(%v) != %v", s, res[i])
		}
	}
}

func TestCheckout(t *testing.T) {
	var cart Cart
	cart.AddItem("eggs")
	cart.AddItem("apple")
	cart.AddItem("milk")
	cart.AddItem("green pepper")
	cart.Checkout() // no output, result print to console
}
