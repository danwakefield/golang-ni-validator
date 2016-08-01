package NI

var disallowedSecond map[string]bool
var disallowedFirstOrSecond map[string]bool
var disallowedCombinations map[string]bool

func init() {
	disallowedFirstOrSecond = map[string]bool{
		"D": true,
		"F": true,
		"I": true,
		"Q": true,
		"U": true,
		"V": true,
	}
	disallowedSecond = map[string]bool{
		"O": true,
	}
	disallowedCombinations = map[string]bool{
		"BG": true,
		"GB": true,
		"KN": true,
		"NK": true,
		"NT": true,
		"TN": true,
		"ZZ": true,
	}
}
func ValidNI(s string) bool {
	l := len(s)
	if l != 9 && l != 8 {
		return false
	}

	if disallowedCombinations[s[0:2]] || disallowedFirstOrSecond[s[0:1]] || disallowedFirstOrSecond[s[1:2]] || disallowedSecond[s[1:2]] {
		return false
	}

	for _, x := range s[3:8] {
		if x < '0' || x > '9' {
			return false
		}
	}

	return l == 8 || (s[8] >= 'A' && s[8] < 'E') || s[8] == ' '
}
