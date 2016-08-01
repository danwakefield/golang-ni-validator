package NI

import (
	"fmt"
	"testing"
)

func TestNI(t *testing.T) {
	fmt.Println(ValidNI("JH080265A"))
	fmt.Println(ValidNI("JH080265 "))
	fmt.Println(ValidNI("JH080265"))
	fmt.Println(ValidNI("BG123412D"))
}

func BenchmarkNI(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ValidNI("JH080265A")
		// ValidNI("JH080265 ")
		// ValidNI("JH080265")
		// ValidNI("BG123412D")
	}
}
