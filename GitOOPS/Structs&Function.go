package main

import (
	"fmt"
	"os"
)

type Trade struct {
	Symbol string
	Volume string
	Price  float64
	Buy    bool
}

func NewTrade(symbol string, volume string, price float64, buy bool) (*Trade, error) {
	if symbol == "" {
		return nil, fmt.Errorf("Symbol can't be empty")
	}

	if volumne <= 0 {
		return nil, fmt.Errorf("Volume can't be negative (was %d", volume)
	}

	if price <= 0 {
		return nil, fmt.Errorf("Price cannot be negative (was %f", price)
	}

	trade := &Trade{
		Symbol: symbol,
		Volume: volume,
		Price:  price,
		Buy:    buy,
	}
	return trade, nil
}

// Value returns the trade value
func (t *Trade) Value() float64 {
	value := float64(t.Volume) * t.Price
	if t.Buy {
		value = -value
	}

	return value
}

func main() {
	t, err := NewTrade("MSFT", 10, 99.98, true)

	if err != nil {
		fmt.Printf("error: cant create trade - % s\n", err)
		os.Exit(1)
	}

	fmt.Println(t.Value())
}
