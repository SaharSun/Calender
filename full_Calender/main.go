package main

import (
	"Fuul_Calender/changing"
	"fmt"
	"time"
)

func about_now() (int, string, int, string, int) {
	t := time.Now()
	
	day_name := t.Weekday().String()
	day_pas := t.YearDay()
	year_now := t.Year()
	month_now := t.Month().String()
	day_now := t.Day()

	return year_now, month_now , day_now, day_name, day_pas
}
func main() {
	var number string
	fmt.Printf("1)What time now \n2)Changing Date \n ~chose:")
	fmt.Scanln(&number)
	switch number {
	case "1":
		y, m, d, dn, dp := about_now()
		jy , jm , jd := changing.Changed_to_jalali(y,2,d)
		fmt.Printf("Today is%s ,the %dnd of %s in %d ",dn,d, m, y)
		fmt.Println()
		fmt.Printf(" -%d days have passed since the beginning of the year.", dp)
		fmt.Println()
		fmt.Printf("IN Jalai:%d/%d/%d ", jy , jm ,jd)

	case "2":
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
