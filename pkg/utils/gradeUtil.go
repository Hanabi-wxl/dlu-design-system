package utils

import "time"

func GetGrade() int64 {
	year := int64(time.Now().Year())
	month := int64(time.Now().Month())
	if month >= 9 {
		return year - 3
	} else {
		return year - 4
	}
}

func GetGradeByYear(year int64) int64 {
	yearNow := int64(time.Now().Year())
	month := int64(time.Now().Month())
	if yearNow == year {
		if month >= 9 {
			return year - 3
		} else {
			return year - 4
		}
	} else {
		return year - 4
	}
}
