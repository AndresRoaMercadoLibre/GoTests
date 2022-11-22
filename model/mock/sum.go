package mock

func Sum(numbers []int) int {
	sum := 0
	for _, numbers := range numbers {
		sum += numbers
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tails := numbers[0:] //colas
			sums = append(sums, Sum(tails))
		}
	}
	return sums
}
