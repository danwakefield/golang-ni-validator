package ni

var disallowedCombinations map[string]bool
var disallowedSecond map[byte]bool
var disallowedFirstOrSecond map[byte]bool
var allowedLast map[byte]bool

func init() {
	disallowedCombinations = map[string]bool{
		"BG": true,
		"GB": true,
		"KN": true,
		"NK": true,
		"NT": true,
		"TN": true,
		"ZZ": true,
	}
	disallowedFirstOrSecond = map[byte]bool{
		'D': true,
		'F': true,
		'I': true,
		'Q': true,
		'U': true,
		'V': true,
	}
	disallowedSecond = map[byte]bool{
		'O': true,
	}
	allowedLast = map[byte]bool{
		'A': true,
		'B': true,
		'C': true,
		'D': true,
		' ': true,
	}

}
func IsValid(s string) bool {
	l := len(s)
	if l != 9 && l != 8 {
		return false
	}

	if disallowedCombinations[s[0:2]] ||
		disallowedFirstOrSecond[s[0]] ||
		disallowedFirstOrSecond[s[1]] ||
		disallowedSecond[s[1]] {
		return false
	}

	for _, x := range s[3:8] {
		if x < '0' || x > '9' {
			return false
		}
	}

	return l == 8 || allowedLast[s[8]]
}
