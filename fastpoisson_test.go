package fastpoisson

import "testing"

func TestPoisson(t *testing.T) {
	fpg := New(0.0005)
	for x := 0; x < 100000; x++ {
		fpg.Poisson()
	}
}
