package numbers

import (
	"testing"
	"Day1/slice"
)


//NOTE TO SELF []int is SLICE, [size]int is ARRAY...

func testGRN(t *testing.T){
	var tests = []struct {
		input []int 
		want []int
	}{
		{[]int{1, 1, 1, 1}, []int{1, 1, 1, 1} },
		{[]int{1, 2, 3, 4}, []int{} },
		{[]int{1, 1, 2, 2}, []int{1, 2} },
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, []int{9} },
		{[]int{7, 7, 7, 2, 3, 7, 7}, []int{7, 7, 7, 7} },
	}

	for _, c := range tests {
		got := GetRepeatingNumbers(c.input, true)

		isWrong := false
		gotLength := len(got)
		wantLength := len(c.want)

		if (gotLength != wantLength ) {
			isWrong = true
		} else {

			//length is same, compare contents
			//only check content, not order! [1 2 3] is same as [3 1 2]
			for _, k := range c.want {

				idx := slice.SliceIndex(gotLength, func(i int) bool { return got[i] == k})

				if (idx < 0){
					isWrong = true
					break
				}
			}
		}

		if (isWrong){
			t.Errorf("GetRepeatingNumbers(%v) == %v, want %v", c.input, got, c.want)
		}

	}
}

func testGRNH(t *testing.T){
	var tests = []struct {
		input []int 
		want []int
	}{
		{[]int{1, 2, 1, 2}, []int{1, 1, 2, 2} },
		{[]int{1, 2, 2, 1}, []int{} },
		{[]int{1, 2, 3, 4, 2, 5}, []int{2, 2} },
		{[]int{1, 2, 3, 1, 2, 3}, []int{1, 2, 3, 1, 2, 3} },
		{[]int{1, 2, 1, 3, 1, 4, 1, 5}, []int{1, 1, 1, 1} },
	}

	for _, c := range tests {
		got := GetRepeatingNumbersHalfway(c.input)

		isWrong := false
		gotLength := len(got)
		wantLength := len(c.want)

		if (gotLength != wantLength ) {
			isWrong = true
		} else {

			//length is same, compare contents
			//only check content, not order! [1 2 3] is same as [3 1 2]
			for _, k := range c.want {

				idx := slice.SliceIndex(gotLength, func(i int) bool { return got[i] == k})

				if (idx < 0){
					isWrong = true
					break
				}
			}
		}

		if (isWrong){
			t.Errorf("GetRepeatingNumbers(%v) == %v, want %v", c.input, got, c.want)
		}

	}
}

func testSS(t *testing.T){
	var tests = []struct {
		input []int 
		want int
	}{
		{[]int{1, 1, 1, 1}, 4},
		{[]int{1, 2, 3, 4}, 10},
		{[]int{1, 1, 2, 2}, 6 },
		{[]int{9, 1, 2, 1, 2, 1, 2, 9}, 27},
	}

	for _, c := range tests {
		got := SumSlice(c.input)

		if (got != c.want ) {

			t.Errorf("SumSlice(%v) == %d, want %d", c.input, got, c.want)
		}
	}	
}


func Test(t *testing.T){
	testGRN(t)
	testGRNH(t)
	testSS(t)
}