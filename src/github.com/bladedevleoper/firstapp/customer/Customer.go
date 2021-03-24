package customer

import (
	"fmt"
	"strconv"
)

var (
	customerName string  = ""
	customerID   int     = 0
	address      string  = ""
	county       string  = ""
	products     string  = ""
	price        float32 = 0
)

func shopping() {

	newLine := "\n"
	customerName = "James Jones"
	customerID = 12345
	address = "8 Bryn lane"
	county = "Blackwood"
	products = "Nike football boots"
	price = 69.99

	//print the outputs in the console
	fmt.Println("customer Id: " + strconv.Itoa(customerID) + newLine)
	fmt.Println("customer name: " + customerName + newLine)
	fmt.Println("customer address: " + address + newLine)
	fmt.Println("county: " + county + newLine)
	fmt.Println("bought: " + products + newLine)
	fmt.Println("price: £" + fmt.Sprintf("%.2f", price) + newLine)
}
