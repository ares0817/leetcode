package problem_20

func isValid(s string) bool {

	a := make([]int32, 0)
	m := map[int32]int32{
		125: 123,
		93: 91,
		41: 40,
	}

	for _, c := range s {
		if left, ok := m[c]; ok {
			if len(a) == 0 {
				return false
			}
			l := len(a) - 1
			if a[l] != left {
				return false
			}
			a = a[:l]
		} else {
			a = append(a, c)
		}
	}

	if len(a) > 0 {
		return false
	}

	return true
}
