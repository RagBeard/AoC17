package numbers

import "testing"


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
	}

	for _, c := range tests {
		got := GetRepeatingNumbers(c.input, true)

		isWrong := false

		if (len(got) != len(c.want) ) {
			isWrong = true
		} else {

			//length is same, compare contents

			for i := 0; i < len(got); i++ {
				if (got[i] != c.want[i]) {
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
	testSS(t)
}