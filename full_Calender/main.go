package main

import (
	"Fuul_Calender/changing"
	"fmt"
	"time"
	"math"
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

func time_sub(start_year, start_month, start_day, end_year, end_month, end_day int) time.Duration {
	// add 0 befor date number in Start-Date
	var start_Date string
	if start_month < 10 {
		start_Date = fmt.Sprintf("%d-0%d-%d", start_year, start_day, start_day)
	} else if start_day < 10 {
		start_Date = fmt.Sprintf("%d-%d-0%d", start_year, start_day, start_day)
	} else if start_month < 10 && start_day < 10 {
		start_Date = fmt.Sprintf("%d-0%d-0%d", start_year, start_day, start_day)
	} else {
		start_Date = fmt.Sprintf("%d-%d-%d", start_year, start_day, start_day)
	}

	// add 0 befor date number in End-Date
	var end_date string
	if end_month < 10 && end_day < 10  {
		end_date = fmt.Sprintf("%d-0%d-0%d", end_year, end_month, end_day)
	} else if end_day < 10 {
		end_date = fmt.Sprintf("%d-%d-0%d", end_year, end_month, end_day)
	} else if end_month < 10  {
		end_date = fmt.Sprintf("%d-0%d-%d", end_year, end_month, end_day)
		
	} else {
		end_date = fmt.Sprintf("%d-%d-%d", end_year, end_month, end_day)
	}

	layout := "2006-01-02"
	date1, _ := time.Parse(layout, start_Date)
	date2, _ := time.Parse(layout, end_date)
	duration := date2.Sub(date1)
	return duration
}


func info_age (duration time.Duration)(hour , year_age ,month_age float64){
	day := math.Floor(duration.Hours()/24)
	fmt.Println( duration , duration.Abs(),day)
	year := day/(30*12)
	fmt.Println(year)
	mm := year * 12
	m := day / 30
	month := math.Abs(m - mm)
	fmt.Println(month)

    return duration.Hours(), math.Floor(year) , math.Floor(month)
}

func main() {
	var number string
	fmt.Printf("1)What time now \n2)Changing Date \n3)How old are you? \n4)Age difference between two people \n ~chose:")
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

	case "3":
		var number_changing_brith string
		fmt.Printf("1)finde in gregorian \n2)finde to jalali \n ~chose:")
		fmt.Scanln(&number_changing_brith)

		switch number_changing_brith {
		case "1":
			var brith_year int
			var birth_month int
			var birth_day int

			fmt.Printf("write year:")
			fmt.Scanln(&brith_year)

			fmt.Printf("write month:")
			fmt.Scanln(&birth_month)

			fmt.Printf("write day:")
			fmt.Scanln(&birth_day)

			yn, mn, dn, _, _ := about_now()
			de := time_sub(brith_year, birth_month, birth_day, yn, month_gregorian[mn], dn)
			fmt.Println(de)

		case "2":
			var brith_year int
			var birth_month int
			var birth_day int

			fmt.Printf("write year:")
			fmt.Scanln(&brith_year)

			fmt.Printf("write month:")
			fmt.Scanln(&birth_month)

			fmt.Printf("write day:")
			fmt.Scanln(&birth_day)

			yn, mn, dn, _, _ := about_now()
			yng, mng, dng := changing.Changed_to_jalali(yn, month_gregorian[mn], dn)

			de := time_sub(brith_year, birth_month, birth_day, yng, mng, dng)
			fmt.Println(de, de.Hours())

		}

	case "4":
		var number_age_difference string
		fmt.Printf("1)Find the age difference of two people in the calendarin gregorian \n2)Find the age difference of two people in the calendar jalali \n ~chose:")
		fmt.Scanln(&number_age_difference)

		switch number_age_difference {
		case "1":
			var brith_year_P1 int
			var birth_month_P1 int
			var birth_day_P1 int
			fmt.Printf("write year of Person 1: ")
			fmt.Scanln(&brith_year_P1)
			fmt.Printf("write month of Person 1: ")
			fmt.Scanln(&birth_month_P1)
			fmt.Printf("write day of Person 1: ")
			fmt.Scanln(&birth_day_P1)

			var brith_year_P2 int
			var birth_month_P2 int
			var birth_day_P2 int
			fmt.Printf("write year of Person 2 : ")
			fmt.Scanln(&brith_year_P2)
			fmt.Printf("write month of Person 2: ")
			fmt.Scanln(&birth_month_P2)
			fmt.Printf("write day of Person 2: ")
			fmt.Scanln(&birth_day_P2)

			de := time_sub(brith_year_P1,birth_month_P1 , birth_day_P1 , brith_year_P2 , birth_month_P1  , birth_day_P2)
			hour , month , year := info_age(de)
			fmt.Println( de ,hour , month , year)

		case "2":
			var brith_year_P1 int
			var birth_month_P1 int
			var birth_day_P1 int
			fmt.Printf("write year of Person 1: ")
			fmt.Scanln(&brith_year_P1)
			fmt.Printf("write month of Person 1: ")
			fmt.Scanln(&birth_month_P1)
			fmt.Printf("write day of Person 1: ")
			fmt.Scanln(&birth_day_P1)

			var brith_year_P2 int
			var birth_month_P2 int
			var birth_day_P2 int
			fmt.Printf("write year of Person 2 : ")
			fmt.Scanln(&brith_year_P2)
			fmt.Printf("write month of Person 2: ")
			fmt.Scanln(&birth_month_P2)
			fmt.Printf("write day of Person 2: ")
			fmt.Scanln(&birth_day_P2)

			ybp1 , mbp1 , dbp1 := changing.Changed_to_gregorian(brith_year_P1, birth_month_P1, birth_day_P1)
			ybp2 , mbp2 , dbp2 := changing.Changed_to_gregorian(brith_year_P2, birth_month_P2, birth_day_P2)
			de := time_sub(ybp1 , mbp1 , dbp1, ybp2 , mbp2 , dbp2)
			fmt.Println(de)

		}

	}

}
