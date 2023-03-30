package helper

func StringInArr(a string, list []string) bool {
	if len(list) == 0 {
		return false
	}

	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}
