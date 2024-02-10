package main

import (
	"Fuul_Calender/changing/changed"
	"fmt"

	
)

func main() {
	var number string
	fmt.Printf("1)changing \n chose:")
	fmt.Scanln(&number)
	switch number {
	case "1":
		e := changed()
		fmt.Println(e)
	}

}
