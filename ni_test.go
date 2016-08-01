package ni

import "testing"

func TestNI(t *testing.T) {
	cases := map[string]bool{
		"JH080265A": true,
		"JH080265 ": true,
		"JH080265":  true,
		"BG123412D": false,
	}

	for in, expected := range cases {
		actual := IsValid(in)
		if actual != expected {
			t.Errorf("IsValid(%s) should return %v, returned %v", in, expected, actual)
		}
	}
}

func BenchmarkNI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsValid("JH080265A")
	}
}
