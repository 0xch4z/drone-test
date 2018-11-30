package math

import "testing"

type Fixture struct {
	In  []int
	Out int
}

func TestAdd(t *testing.T) {
	for _, fixture := range []Fixture{
		{[]int{3, 3}, 6},
		{[]int{4, 12}, 15},
	} {
		if res := Add(fixture.In[0], fixture.In[1]); res != fixture.Out {
			t.Errorf("Add(%v); expected %v, got %v",
				fixture.In, fixture.Out, res)
		}
	}
}
