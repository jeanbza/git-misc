package main

import "testing"

func TestOrderCoffee(t *testing.T) {
    const in, out = 3, 9
    if x := OrderCoffee(in); x != out {
        t.Errorf("OrderCoffee(%v) = %v, want %v", in, x, out)
    }
}