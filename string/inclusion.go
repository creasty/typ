package string_h

func Include(item string, ary []string) bool {
	for _, v := range ary {
		if v == item {
			return true
		}
	}

	return false
}
