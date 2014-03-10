// Package provide an implementation of the epsilon-Greedy bandit algorithm.
package epsilongreedy



import "errors"
import "math/rand"



type T struct {
	size          uint64
	epsilon       float64
	indexForMax   uint64
	valueForMax   float64
	counts      []uint64  //@TODO: How does this handle overflow?
	values      []float64 //@TODO: How does this handle overflow?
}



func New(size uint64, epsilon float64) *T {

	// Don't allow size to be larger than 2^63 (= 9,223,372,036,854,775,808)
	if 1 > size || 9223372036854775808 < size {
		return nil
	}

	// Epsilon must be between 0 and 1 inclusive.
	// (Although it likely be a non-zero but closer to zero than one.)
	if epsilon > 1 || epsilon < 0 {
		return nil
	}


	me := T{
		size        : size,
		epsilon     : epsilon,
		indexForMax : 0, //@TODO: Should this be randomized?
		valueForMax : 0.0,
		counts      : make([]uint64, size, size),
		values      : make([]float64, size, size),
	}


	return &me
}



func (me *T) exploit() uint64 {
	return me.indexForMax
}



func (me *T) explore() uint64 {
	return uint64(rand.Int63n(int64(me.size)-1))
}



func (me *T) Select() uint64 {

	// @TODO: Should I be worrying about the distribution of this pseudo-random number generator?
	// @TODO: Is this a problem that it returns from [0,1) instead of [0,1]
	r := rand.Float64()

	return selectIt(r, me.epsilon, me.exploit, me.explore)
}

type selector func() uint64

func selectIt(r float64, epsilon float64, exploit selector, explore selector) uint64 {

	if epsilon < r {
		return exploit()
	} else {
		return explore()
	}
}



func (me *T) Update(index uint64, reward float64) error {

	if me.size <= index {
		// Error!
		return errors.New("index out of range")
	}


	n := me.counts[index]
	n++
	me.counts[index] = n
	
	n_inverse := 1.0/float64(n)


	old_value := me.values[index]
	new_value := float64(n-1)*n_inverse*old_value + n_inverse*reward
	me.values[index] = new_value


	// See if there is a new "best".
	if new_value > me.valueForMax {
		me.valueForMax = new_value
		me.indexForMax = index
	}


	return nil
}
