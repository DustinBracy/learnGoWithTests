package arrays

// Sum returns the sum of a slice of integers.
func Sum(numbers []int) int {
	add := func(acc, x int) int {
		return acc + x
	}
	return Reduce(numbers, add)
	// sum := 0
	// for _, number := range numbers {
	// 	sum += number
	// }
	// return sum
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
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

func Reduce[T any](collection []T, reducer func(T, T) T) T {
	if len(collection) == 0 {
		var zeroValue T
		return zeroValue
	}
	var value T

	for i, item := range collection {
		if i == 0 {
			value = item
			continue
		}

		value = reducer(value, collection[i])
	}
	return value

}
