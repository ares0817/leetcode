package problem_17

func letterCombinations(digits string) []string {

	m := map[int32][]string{
		50: []string{"a", "b", "c"},
		51: []string{"d", "e", "f"},
		52: []string{"g", "h", "i"},
		53: []string{"j", "k", "l"},
		54: []string{"m", "n", "o"},
		55: []string{"p", "q", "r", "s"},
		56: []string{"t", "u", "v"},
		57: []string{"w", "x", "y", "z"},
	}

	a1 := make([]string, 0)

	for _, c := range digits {
		a2 := make([]string, 0)
		arr, exist := m[c]
		if !exist {
			continue
		}
		if len(a1) == 0 {
			a1 = arr
			continue
		}

		for _, s1 := range a1 {
			for _, s2 := range arr {
				a2 = append(a2, s1 + s2)
			}
		}
		a1 = a2
	}

	return a1
}
