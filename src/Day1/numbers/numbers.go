package numbers

func GetRepeatingNumbers(input []int, circular bool) []int {

	//make it "circular" by appending first element last
	if (circular){
		input = append(input, input[0])
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

func SumSlice(slice []int) int {
	total := 0

	for _, c := range slice {
		total += c
	}

	return total
}