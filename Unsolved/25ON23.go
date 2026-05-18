// 2025 Oct Nov Paper 23 Question 10
// Question is player 1 and player 2 rolls 100 time a number between 1-6
// All of these numbers are stored in their respective arrays
// if player 1 > player 2 at the moment, player 1 gets 2 points and vice versa
// if they match then they get 1 point each
// at the end of the game if the points match then they roll 1 more time
// player with the higher amount of points is output as winner

package on2523

import (
	"fmt"
	"math/rand"
)

func ON2523() {

	fmt.Println("Hello World")
	playerRoll := playerNoGenerator()

}

func playerNoGenerator() [][]int {
	var numberGenerated [][]int

	for k := 1; k <= 2; k++ {
		for i := 1; i < 100; i++ {
			numberGenerated[k][i] = rand.Intn(6)
		}
	}
	return numberGenerated
}
