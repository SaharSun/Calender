package main

import (
	"Fuul_Calender/changing"
	"fmt"
)

func main() {
	var number string
	fmt.Printf("1)changing \n ~chose:")
	fmt.Scanln(&number)
	switch number {
	case "1":
		var number_changing string
		fmt.Printf("1)changing to gregorian \n2)changing to jalali \n ~chose:")
		fmt.Scanln(&number_changing)
		switch number_changing {
		case "1":
			var year_change int
			var month_change int
			var day_change int

			fmt.Printf("write year:")
			fmt.Scanln(&year_change)

			fmt.Printf("write month:")
			fmt.Scanln(&month_change)

			fmt.Printf("write day:")
			fmt.Scanln(&day_change)

			e, r, t := changing.Changed_to_gregorian(year_change, month_change, day_change)
			fmt.Println(e, r, t)
		case "2":
			var year_change int
			var month_change int
			var day_change int

			fmt.Printf("write year:")
			fmt.Scanln(&year_change)

			fmt.Printf("write month:")
			fmt.Scanln(&month_change)

			fmt.Printf("write day:")
			fmt.Scanln(&day_change)

			e, r, t := changing.Changed_to_jalali(year_change, month_change, day_change)
			fmt.Println(e, r, t)
		}

	}

}
