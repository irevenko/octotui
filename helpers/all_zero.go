package helpers

func AllZero(s []float64) bool {
	for _, v := range s {
		if v != 0 {
			return false
		}
	}
	return true
}
