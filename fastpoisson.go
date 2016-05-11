package fastpoisson

import (
	"math"
	"math/rand"
)

type FastPoissonGenerator struct {
	L float64
	R *rand.Rand
}

/* You need to supply your own Rand instance to avoid locking which would slow us down. No concurrent use. */
func New(lambda float64, rng *rand.Rand) *FastPoissonGenerator {
	return &FastPoissonGenerator{L: math.Pow(math.E, -lambda), R: rng}
}

/*
   Minimal Knuth Algorithm from
   https://en.wikipedia.org/wiki/Poisson_distribution#Generating_Poisson-distributed_random_variables
   with cached L. For lambda < 1 this was a 4x speed up over the fastest implementation for Go we found.

   If you need further speed for large lambdas the table based algorithm from
     https://www.jstatsoft.org/article/view/v011b03/v11b03.pdf

   should beat almost anything. There's a matlab implementation here:
   http://www.mathworks.com/matlabcentral/fileexchange/7309-randraw
*/

func (fpg *FastPoissonGenerator) Poisson() int64 {
	var k int64 = 0
	var p float64 = 1.0

	for p > fpg.L {
		k++
		p *= fpg.R.Float64()
	}
	return (k - 1)
}
