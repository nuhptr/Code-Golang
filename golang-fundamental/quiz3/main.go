package main

import "fmt"

func main() {
	/*scores := [8]int{
		100,80,75,92,70,93,88,67,
	}
	var total int

	for _, value := range scores {
		total = total + value
	}

	length := len(scores)
	rate := float64(total) / float64(length)
	fmt.Println(rate)*/

	scores := []int {
		100, 80, 75, 92, 70, 93, 88, 67,
	}
	var goodScores []int

	for _, score := range scores{
		if score >= 90{
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores) // [100, 92, 93]
}
