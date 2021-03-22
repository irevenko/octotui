package helpers

func UserDataBound(slice []string) int {
	if len(slice) > 4 { // 4 is max amount of entries for pie/bar chart
		return 4
	} else if len(slice) < 1 {
		return 0
	} else if len(slice) < 2 {
		return 1
	} else if len(slice) < 3 {
		return 2
	} else if len(slice) < 4 {
		return 3
	} else if len(slice) < 5 {
		return 4
	}

	return 0
}

func OrgDataBound(slice []string) int {
	if len(slice) > 7 { // 7 is max amount of entries for pie/bar chart
		return 7
	} else if len(slice) < 1 {
		return 0
	} else if len(slice) < 2 {
		return 1
	} else if len(slice) < 3 {
		return 2
	} else if len(slice) < 4 {
		return 3
	} else if len(slice) < 5 {
		return 4
	} else if len(slice) < 6 {
		return 5
	} else if len(slice) < 7 {
		return 6
	} else if len(slice) < 8 {
		return 7
	}

	return 0
}


