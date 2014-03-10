package epsilongreedy



import "testing"
import "math/rand"



func TestSelect(t *testing.T) {

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


		index := x.Select()

		if 0 > index || size <= index {
			t.Errorf("Selected index is not the expected range for size [%+v] and epsilon [%+v] and object [%+v]", size, epsilon, x)
		}
	}

}



func TestSelectIt(t *testing.T) {

	tests := []struct{
		size    uint64
		epsilon float64


		series []struct{
			r float64

			exploitValue uint64
			exploreValue uint64

			expectedIndex uint64
		}
	}{
		{
			10,
			0.1,

			[]struct{
				r float64

				exploitValue uint64
				exploreValue uint64

				expectedIndex uint64
			}{
				{
					0.2,

					3,
					7,

					3,
				},
				{
					0.04,

					3,
					7,

					7,
				},
				{
					0.5,

					2,
					4,

					2,
				},
				{
					0.005,

					2,
					4,

					4,
				},
			},
		},
	}



	for _,datum := range tests {

		x := New(datum.size, datum.epsilon)

		if nil == x {
			t.Errorf("Did not expect a nil return for test datum [%+v]", datum)
		}

		if nil == x {
			continue
		}


		for _,seriesItem := range datum.series {

			exploitValue := seriesItem.exploitValue
			exploreValue := seriesItem.exploreValue

			actualIndex := selectIt(seriesItem.r, datum.epsilon, func() uint64 {return exploitValue}, func() uint64 {return exploreValue})

			if seriesItem.expectedIndex != actualIndex {
				t.Errorf("Actual selected index [%+v] is different than expected index [%+v].", seriesItem.expectedIndex, actualIndex);
			}
		}


	}
}
