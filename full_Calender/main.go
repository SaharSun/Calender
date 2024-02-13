package main

import (
	"Fuul_Calender/changing"
	"fmt"
	"time"
)

var month_gregorian = map[string]int{
	"January":   1,
	"February":  2,
	"March":     3,
	"April":     4,
	"May":       5,
	"June":      6,
	"July":      7,
	"August":    8,
	"September": 9,
	"October":   10,
	"November":  11,
	"December":  12,
}

var month_jalali = map[int]string{
	1:  "Farvardin",
	2:  "Ordibehesht",
	3:  "Khordad",
	4:  "Tir",
	5:  "Mordad",
	6:  "Shahrivar",
	7:  "Mehr",
	8:  "Aban",
	9:  "Azar",
	10: "Dey",
	11: "Bahman",
	12: "Esfand",
}

var day_name_jalali = map[string]string{
	"Saturday":  "Shanbe",
	"Sunday":    "Yek-Shanbe",
	"Monday":    "Do-Shanbe",
	"Tuesday":   "Se-Shanbe",
	"Wednesday": "Chahar-Shanbe",
	"Thursday":  "Panj-Shanbe",
	"Friday":    "Jomeae",
}

func about_now() (int, string, int, string, int) {
	t := time.Now()

	day_name := t.Weekday().String()
	day_pas := t.YearDay()
	year_now := t.Year()
	month_now := t.Month().String()
	day_now := t.Day()

	return year_now, month_now, day_now, day_name, day_pas
}

// func time_sub(start , end string)int{
// 	layout := "2006-01-02"
//     date1, _ := time.Parse(layout, start)
//     date2, _ := time.Parse(layout, end)

//     duration := date2.Sub(date1)
// 	d := duration.Hours()/24
//     fmt.Println( duration , d/365 ,duration.Hours() )
// }
func main() {
	var number string
	fmt.Printf("1)What time now \n2)Changing Date \n3)How old are you? \n ~chose:")
	fmt.Scanln(&number)
	switch number {
	case "1":
		y, m, d, dn, dp := about_now()
		jy, jm, jd := changing.Changed_to_jalali(y, 2, d)
		fmt.Println()
		fmt.Printf("Today is '%s' ,the %dnd of '%s' in %d -> (%d / %d / %d )", dn, d, m, y, y, month_gregorian[m], d)
		fmt.Println()
		fmt.Printf(" ~ %d days have passed since the beginning of the year.", dp)
		fmt.Println()
		fmt.Println()
		fmt.Println("IN Jalai:")
		fmt.Printf("Today is '%s' ,the %dnd of '%s' in %d -> (%d / %d / %d )", day_name_jalali[dn], jd, month_jalali[jm], jy, jy, jm, jd)
		fmt.Println()

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
			fmt.Printf("%d / %d / %d", e, r, t)
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
			fmt.Printf("%d / %d / %d", e, r, t)
		}

	case "3" :
		var year_old int
		var month_old int
		var day_old int

		fmt.Printf("write year:")
		fmt.Scanln(&year_old)

		fmt.Printf("write month:")
		fmt.Scanln(&month_old)

		fmt.Printf("write day:")
		fmt.Scanln(&day_old)

		layout := "2006-01-02"
		y, m, d, _, _ := about_now()
		birthday := fmt.Sprintf("%d-%d-%d", year_old , month_old , day_old)
		now_date := fmt.Sprintf("%d-%d-%d",y , m ,d )
		date1, _ := time.Parse(layout,birthday)
    	date2, _ := time.Parse(layout, now_date)
		duration := date2.Sub(date1)
		fmt.Println( duration.Hours())

	}

}
