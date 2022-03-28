package calculation

func Hitung(number int, numberTwo int, numberThree int) int {
	hasil := number + numberTwo + numberThree
	scoreHasil := hasil / 3
	return scoreHasil
}

func Score(number int) string {
	if number >= 80 {
		return "A"
	} else if number > 70 && number < 79 {
		return "B"
	} else if number > 60 && number < 69 {
		return "C"
	} else if number > 50 && number < 59 {
		return "D"
	} else {
		return "E"
	}
}
