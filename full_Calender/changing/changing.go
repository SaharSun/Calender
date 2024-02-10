package changing

// import "fmt"

type date struct {
	year  int
	month int
	day   int
}

func (d *date) gregorian_to_jalali() (int, int, int) {
	var jy, jm, jd, gy2, days int
	var g_d_m = [12]int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334}
	if d.month > 2 {
		gy2 = d.year + 1
	} else {
		gy2 = d.year
	}
	days = 355666 + (365 * d.year) + ((gy2 + 3) / 4) - ((gy2 + 99) / 100) + ((gy2 + 399) / 400) + d.day + g_d_m[d.month-1]
	jy = -1595 + (33 * (days / 12053))
	days %= 12053
	jy += 4 * (days / 1461)
	days %= 1461
	if days > 365 {
		jy += (days - 1) / 365
		days = (days - 1) % 365
	}
	if days < 186 {
		jm = 1 + (days / 31)
		jd = 1 + (days % 31)
	} else {
		jm = 7 + ((days - 186) / 30)
		jd = 1 + ((days - 186) % 30)
	}
	return jy, jm, jd
}
func changed() (int, int, int) {
	w := &date{
		year:  2024,
		month: 2,
		day:   6,
	}

	jy, jm, jd := w.gregorian_to_jalali()
	return jy, jm, jd

}
