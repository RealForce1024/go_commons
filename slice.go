package go_commons

func compare(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}

	//for: []int{} != []int(nil)
	if (s1 == nil) != (s2 == nil) {
		return false
	}

	//for:bounds check
	s2 = s2[:len(s1)]

	for i, v := range s1 {
		if v != s2[i] {
			return false
		}
	}
	return true
}
