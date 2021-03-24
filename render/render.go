package render

import (
	"fmt"
	"strconv"
)

//struct literal
type person struct {
	name       string
	cardNumber string
}

func RenderShoppingCart(name string, payload []string) {

	fmt.Println("Hello, " + name + " your shopping list today is")
	//item start
	item := 1
	var customer person
	customer.cardNumber = "111112111"
	for i := 0; i < len(payload); i++ {

		fmt.Println("item:" + strconv.Itoa(item) + " " + payload[i])
		//item iterate
		item++
	}

	fmt.Println(name + " has payed with the following card " + customer.cardNumber)

}
