package numbers


func GetRepeatingNumbers(input []int, circular bool) []int {

	//make it "circular" by appending all repeating elements in the beginning,
	// to the end, if it matches the last element
	if (circular){

		//pick first element from the start, and any following of the same value
		//remove them from there and append to end of array	
		start := []int{input[0]}
		var nStartIndex int = 0

		for i := 1; i < len(input); i++ {

			if (input[i] == start[0]) {
				start = append(start, input[i])
			} else {
				nStartIndex = i
				break
			}
			if (i == len(input) - 1){
				nStartIndex = i
			}
		}

		//trim off the start chunk
		input = input[nStartIndex:]

		//add it to the end
		input = append(input, start...)

	}

	result := make([]int, 0)
	length := len(input)

	for i := 0; i < length; i++ {

		if (i+1 >= length)	{
			break
		}

		if (input[i] == input[i+1]) {
			result = append(result, input[i])
		}
	}


 	return result
}


func GetRepeatingNumbersHalfway(input []int) []int {

	result := make([]int, 0)
	length := len(input)
	halfway := length/2

	for i := 0; i < length; i++ {

		halfwayIdx := (i + halfway)%length

		if (input[i] == input[halfwayIdx]) {
			result = append(result, input[i])
		}
	}


 	return result
}

func SumSlice(slice []int) int {
	total := 0

	for _, c := range slice {
		total += c
	}

	return total
}