// 2025 Oct Nov Paper 22 Question 8
package on2522

import (
	"fmt"
	"math/rand"
)

func ON2522() {
	freq := generate()
	fmt.Println("Below is the freq for each number: ")
	for _, pair := range freq {
		fmt.Printf("Number %d: is shown %d times.\n", pair[0], pair[1])
	}

	//copying freq values to new slice
	sort := make([][]int, len(freq))
	copy(sort, freq)

	totals := len(sort)

	//sorting algo [Selection sort]
	for i := 0; i < totals-1; i++ {
		top := i
		for j := i + 1; j < totals; j++ {
			if sort[j][1] > sort[top][1] {
				top = j
			}
		}
		sort[i], sort[top] = sort[top], sort[i]
	}

	//Output the freq sorts
	fmt.Println("\nFrequencies arranged in ascending order: ")
	for _, pair := range sort {
		fmt.Printf("Number %d is shown %d times\n", pair[0], pair[1])
	}
}

func generate() [][]int {
	// init the slice
	var goodName [][]int

	//init this slice to ([1,0],[2,0],[3,0],...) for 10 numbers
	for i := 0; i < 10; i++ {
		goodName = append(goodName, []int{i, 0})
	}

	//generate the random numbers
	for j := 1; j <= 100000; j++ {
		randomNo := rand.Intn(10) + 1
		goodName[randomNo-1][1]++
	}
	return goodName
}
