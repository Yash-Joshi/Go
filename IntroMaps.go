package main

import (
	"fmt"
)

func main() {

	stocks := map[string]float64{
		"AMZN": 169.99,
		"GOOG": 225.98,
		"MSFT": 669.515,
	}

	//Number of Items
	fmt.Println(len(stocks))

	//Get a Value
	fmt.Println(stocks["MSFT"])

	//Get Zero value if not found
	fmt.Println(stocks["TESLA"])

	//Use two value to see if found
	value, ok := stocks["TESLA"]

	if !ok {
		fmt.Println("Not found stock")
	} else {
		fmt.Println(value)
	}

	//SET
	stocks["TESLA"] = 32.22
	fmt.Println(stocks)

	//Delete
	delete(stocks, "MSFT")
	fmt.Println(stocks)

	//Single key for is on keys
	for key := range stocks {
		fmt.Println(key)
	}

	//Double value "for" is key , value
	for key, value := range stocks {
		fmt.Println("%s --> %.2f\n", key, value)
	}
}
