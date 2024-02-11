package changing

// import "fmt"

type date struct {
	year  int
	month int
	day   int
}

// changing to jalali
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
func Changed_to_jalali(y, m, d int) (int, int, int) {
	w := &date{
		year:  y,
		month: m,
		day:   d,
	}

	jy, jm, jd := w.gregorian_to_jalali()
	return jy, jm, jd
}

// changing to gregorian
func (d *date) jalali_to_gregorian() (int, int, int) {
	var gy, gm, gd, days int
	var sal_a = [13]int{0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	d.year += 1595
	days = -355668 + (365 * d.year) + ((d.year / 33) * 8) + (((d.year % 33) + 3) / 4) + d.day
	if d.month < 7 {
		days += (d.month - 1) * 31
	} else {
		days += ((d.month - 7) * 30) + 186
	}
	gy = 400 * (days / 146097)
	days %= 146097
	if days > 36524 {
		days--
		gy += 100 * (days / 36524)
		days %= 36524
		if days >= 365 {
			days++
		}
	}
	gy += 4 * (days / 1461)
	days %= 1461
	if days > 365 {
		gy += (days - 1) / 365
		days = (days - 1) % 365
	}
	gd = days + 1
	if (gy%4 == 0 && gy%100 != 0) || (gy%400 == 0) {
		sal_a[2] = 29
	}
	gm = 0
	for gm < 13 && gd > sal_a[gm] {
		gd -= sal_a[gm]
		gm++
	}
	return gy, gm, gd
}

func Changed_to_gregorian(y, m, d int) (int, int, int) {
	w := &date{
		year:  y,
		month: m,
		day:   d,
	}

	jy, jm, jd := w.jalali_to_gregorian()
	return jy, jm, jd
}
