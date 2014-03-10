package epsilongreedy



import "testing"
import "math/rand"



func TestExplore(t *testing.T) {

	for i:=0; i<100; i++ {

		size := 1+uint64(rand.Int63n(1000))

		// rand.Float64() returns [0,1)
		epsilon := rand.Float64()
		
		x := New(size, epsilon)

		if nil == x {
			t.Errorf("Expected a non-nil return for size [%+v] and epsilon [%v]", size, epsilon)
		}

		if epsilon != x.epsilon {
			t.Errorf("Stored epsilon and test epsilon are different for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}


		index := x.explore()

		if 0 > index || size <= index {
			t.Errorf("Selected index is not the expected range for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}

		if 0 > index || size <= index {
			t.Errorf("The returned index for exploring [%+v] is not the interval [0,%v).", index, size)
		}

	}

}
