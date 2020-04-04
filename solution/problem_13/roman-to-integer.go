package problem_13

func romanToInt(s string) int {

	idx := 0
	l := len(s)

	res := 0

	for idx < l {
		switch s[idx] {
		case 77:
			//M
			res += 1000
			idx++
		case 68:
			//D
			res += 500
			idx++
		case 67:
			//C
			if idx + 1 < l && s[idx + 1] == 77 {
				res += 900
				idx += 2
			} else if idx + 1 < l && s[idx + 1] == 68 {
				res += 400
				idx += 2
			} else {
				res += 100
				idx++
			}
 		case 76:
			//L
			res += 50
			idx++
		case 88:
			//X
			if idx + 1 < l && s[idx + 1] == 67 {
				res += 90
				idx += 2
			} else if idx + 1 < l && s[idx + 1] == 76 {
				res += 40
				idx += 2
			} else {
				res += 10
				idx++
			}
		case 86:
			//V
			res += 5
			idx++
		case 73:
			//I
			if idx + 1 < l && s[idx + 1] == 88 {
				res += 9
				idx += 2
			} else if idx + 1 < l && s[idx + 1] == 86 {
				res += 4
				idx += 2
			} else {
				res += 1
				idx++
			}
		default:
			return res
		}
	}

	return res
}