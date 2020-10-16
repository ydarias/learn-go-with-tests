package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersCollections ...[]int) []int {
	var sums []int

	for _, collection := range numbersCollections {
		sums = append(sums, Sum(collection))
	}

	return sums
}

func SumAllTails(numbersCollections ...[]int) []int {
	var tails []int

	for _, collection := range numbersCollections {
		tails = append(tails, Sum(getTail(collection)))
	}

	return tails
}

func getTail(numbersCollection []int) []int {
	if len(numbersCollection) == 0 {
		return []int{0}
	} else {
		return numbersCollection[1:]
	}
}
