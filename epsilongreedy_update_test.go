package epsilongreedy



import "testing"
import "math/rand"



func TestUpdate(t *testing.T) {

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


		beforeCount := x.counts[index]
		beforeValue := x.values[index]

		expectedCount := beforeCount + 1
		
		reward := 10*rand.Float64()

		x.Update(index, reward)

		if actualCount := x.counts[index] ; expectedCount != actualCount {
			t.Errorf("Count for index [%+v] is not as expected. Expected [%+v] but actual is [%+v]. For size [%+v] and epsilon [%+v] and object [%+v]", index, expectedCount, actualCount, size, epsilon, x)
		}

		if actualValue := x.values[index] ; beforeValue > actualValue {
			t.Errorf("Value for index [%+v] is less than expected. Expected to be greater-than-or-equal-to [%+v] but actual is [%+v]. For size [%+v] and epsilon [%+v] and object [%+v]", index, beforeValue, actualValue, size, epsilon, x)
		}
	}

}
