//each go must have this at the top
package main

import (
	"fmt"
	"strconv"

	"github.com/bladedevleoper/firstapp/customer"

)

//var person string = "James"

//var block to make things more readable instead of redeclaring var each time
//which we can show that these variables are related
var (
	actorName    string = "James Jones"
	companion    string = "Laura Jones"
	doctorNumber int    = 3
	Season       int    = 11
)

func main() {
	//fmt.Println(actorName)
	var j int = 55
	//j = float32(doctorNumber)
	//var actorName string = "Rebecca Jones"
	//var name string = "james"
	var i string = strconv.Itoa(j) //Integer to ascii string
	//i := 42

	//name = "James"
	//fmt.Println(number)
	//fmt.Println(name)
	fmt.Printf("%v, %T \n", i, i)
	//fmt.Println()
	customer.Shopping()
}