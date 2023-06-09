package v3

func size(num int, unit int) (n, rem int) {
	return num / unit, num - (num/unit)*unit
}