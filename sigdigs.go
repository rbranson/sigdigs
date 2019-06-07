package sigdigs

// CountSignificantDigits returns the number of significant digits in the
// passed string. Valid input formats are integers, decimals, and e-notation.
// Passing invalid strings will result in undefined behavior.
func CountSignificantDigits(num string) int {
	digits := 0
	zcount := 0
	counting := false
	dec := false

OUTER:
	for _, c := range num {
		switch c {
		case 'e':
			fallthrough
		case 'E':
			break OUTER
		case '+':
		case '-':
		case '.':
			dec = true
		case '0':
			if counting {
				digits++
			}
			zcount++
		default:
			if !counting {
				counting = true
			}
			digits++
			zcount = 0
		}
	}

	if dec {
		return digits
	}
	return digits - zcount
}
