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
	//init player 1 and 2 totals
	player1Total := 0
	player2Total := 0

	//run the numberGenerator function
	playerRoll := playerNoGenerator()

	//display the numbers for each player. total 100 times each
	for i := 0; i < 100; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Player %d has number: %d \n", j+1, playerRoll[j][i])
		}

		//compares the number for player 1 and 2
		if playerRoll[0][i] > playerRoll[1][i] { //if player 1 wins +2 points to total score
			player1Total = player1Total + 2
		} else if playerRoll[0][i] < playerRoll[1][i] { //if player 2 wins +2 points to total score
			player2Total = player2Total + 2
		} else { //if its a tie then +1 to player 1 and 2 totals
			player1Total = player1Total + 1
			player2Total = player2Total + 1
		}
	}

	//if tie it just outputs the score for 1 of them then rolls 1 number as sudden death
	for {
		fmt.Printf("Its a Tie with %d points", player1Total)
		fmt.Printf("Sudden death")
		player1SD := rand.Intn(6) + 1
		player2SD := rand.Intn(6) + 1

		if player1SD > player2SD {
			fmt.Printf("Player 1 wins with number : %d\n", player1SD)
			break
		} else if player1SD < player2SD {
			fmt.Printf("Player 2 wins with number : %d\n", player2SD)
			break
		}
		fmt.Printf("Sudden death was a tie : %d. Rolling again", player1SD)
	}
	//condition to output who is the winner
	//outputs both winner and loser points too

	if player1Total > player2Total {
		fmt.Printf("Player 1 Wins with %d points\n", player1Total)
		fmt.Printf("Player 2 Lost with %d points", player2Total)
	} else if player2Total > player1Total {
		fmt.Printf("Player 2 Wins with %d points\n", player2Total)
		fmt.Printf("Player 1 Lost with %d points", player1Total)
	}

}

func playerNoGenerator() [2][100]int {
	// declare a 2D array for 2 players and 100 numbers
	var numberGenerated [2][100]int

	for k := 0; k < 2; k++ {
		for i := 0; i < 100; i++ {
			numberGenerated[k][i] = rand.Intn(6) + 1
		}
	}
	return numberGenerated
}
