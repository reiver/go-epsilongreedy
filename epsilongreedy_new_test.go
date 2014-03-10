package epsilongreedy



import "testing"
import "math/rand"



func TestNew(t *testing.T) {

	tests := []struct{
		size                uint64
		epsilon             float64
		expectedReturnIsNil bool
	}{
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 0,
			expectedReturnIsNil: false,
		},
		{
			size:    0,
			epsilon: 0,
			expectedReturnIsNil: true,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 1,
			expectedReturnIsNil: false,
		},
		{
			size:    0,
			epsilon: 1,
			expectedReturnIsNil: true,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 0.5,
			expectedReturnIsNil: false,
		},
		{
			size:    0,
			epsilon: 0.5,
			expectedReturnIsNil: true,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 0.9,
			expectedReturnIsNil: false,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 0.999999,
			expectedReturnIsNil: false,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 0.0000001,
			expectedReturnIsNil: false,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 20,
			expectedReturnIsNil: true,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: -3,
			expectedReturnIsNil: true,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: -1000,
			expectedReturnIsNil: true,
		},
		{
			size:    uint64(1+rand.Int63n(1000)),
			epsilon: 1 + 9223372036854775808, // = 1 + 2^63
			expectedReturnIsNil: true,
		},
	}


	
	for _,datum := range tests {

		x := New(datum.size, datum.epsilon)

		if datum.expectedReturnIsNil && nil != x {
			t.Errorf("Expected a nil return for test datum [%+v]", datum)
		}

		if nil == x {
			continue
		}

		if datum.epsilon != x.epsilon {
			t.Errorf("Stored epsilon and test epsilon are different for test datum [%+v] and object [%+v]", datum, x)
		}

		if datum.size != x.size {
			t.Errorf("Stored size and test size are different for test datum [%+v] and object [%+v]", datum, x)
		}

		if datum.size != uint64(len(x.counts)) {
			t.Errorf("Length of stored counts and test size are different for test datum [%+v] and object [%+v]", datum, x)
		}

		if datum.size != uint64(len(x.values)) {
			t.Errorf("Length of stored values and test size are different for test datum [%+v] and object [%+v]", datum, x)
		}
	}

}



func TestNewWithRandEpsilon(t *testing.T) {

	for i:=0; i<100; i++ {

		size := 1+uint64(rand.Int63n(1000))

		// rand.Float64() returns [0,1)
		epsilon := rand.Float64()
		
		x := New(size, epsilon)

		if nil == x {
			t.Errorf("Expected a non-nil return for epsilon [%v]", epsilon)
		}

		if epsilon != x.epsilon {
			t.Errorf("Stored epsilon and test epsilon are different for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}

		if size != x.size {
			t.Errorf("Stored size and test size are different for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}

		if size != uint64(len(x.counts)) {
			t.Errorf("Length of stored counts and test size are different for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}

		if size != uint64(len(x.values)) {
			t.Errorf("Length of stored values and test size are different for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}
	}


	for i:=0; i<100; i++ {

		size := 1+uint64(rand.Int63n(1000))

		// rand.Float64() returns [0,1)
		epsilon := 1 + rand.Float64()
		if 0 == rand.Int31n(1) {
			epsilon += rand.Float64()
		}
		if 0 == rand.Int31n(1) {
			epsilon *= 1 + rand.Float64()
		}
		if 0 == rand.Int31n(1) {
			epsilon *= -1
		}
		
		x := New(size, epsilon)

		if nil != x {
			t.Errorf("Expected a nil return for size [%+v] and epsilon [%v]", size, epsilon)
		}
	}

}
