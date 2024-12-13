package myMath

func ceil(elem float32) int {
	intPart := int(elem)
	if elem-float32(intPart) == 0 {
		return intPart
	}
	return intPart + 1
}
