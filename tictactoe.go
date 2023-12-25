package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func is_legal(board [3][3]string, x int, y int) bool {
	// out of bound
	if x < 0 || x > 2 || y < 0 || y > 2 {
		return false
	}

	// already filled boxes
	if board[x][y] == "X" || board[x][y] == "Y" {
		return false
	}
	return true
}

func get_winner_or_nil(board [3][3]string) interface{} {
	winner := "-"

	if board[0][0] == board[0][1] && board[0][0] == board[0][2] { // check horizontal
		winner = board[0][0]
	} else if board[1][0] == board[1][1] && board[1][0] == board[1][2] {
		winner = board[1][0]
	} else if board[2][0] == board[2][1] && board[2][0] == board[2][2] {
		winner = board[2][0]
	} else if board[0][0] == board[1][0] && board[0][0] == board[2][0] { // check vertical
		winner = board[0][0]
	} else if board[0][1] == board[1][1] && board[0][1] == board[2][1] {
		winner = board[0][1]
	} else if board[0][2] == board[1][2] && board[0][2] == board[2][2] {
		winner = board[0][2]
	} else if board[0][0] == board[1][1] && board[0][0] == board[2][2] { // diagonal
		winner = board[0][0]
	}

	if winner == "0" || winner == "X" {
		return winner
	}

	return nil
}

func print_board(board [3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Printf("  %s\t", board[i][j])
		}
		fmt.Println("\n---------------------")
	}
}

func main() {
	var (
		board        [3][3]string
		p1           string
		p2           string
		current_turn string
		player_input string
	)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Rules:\n" +
		"1. Box filled with - are empty boxes. Player need to input indexes of box " +
		"2. To fill it with their letter (0 or X).\n" +
		"\tFor e.g.\n" +
		"\ta. Indexes starts from 0\n" +
		"\tb. to fill the first box press: 0 0\n" +
		"\tc. to fill the middle box press: 1 1\n" +
		"\td. to fill the top right box press: 0 2 and so on\n" +
		"3. Invalid input such as 3 3 will lead to losing the game.\n" +
		"4. Also trying to fill a box that already have 0 or X will lead to the other player winning\n" +
		"----------------------------------------------------------------------------------------------------------")

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = "-"
		}
	}

	fmt.Print("Enter 1st Player (0 or X): ")
	_, err := fmt.Scan(&p1)
	if err != nil {
		return
	}

	if p1 == "X" {
		p2 = "0"
	} else {
		p2 = "X"
	}
	fmt.Println("First Player:", p1, ", Second Player:", p2)

	current_turn = p1

	for turns := 1; turns <= 9; turns++ {
		print_board(board)

		fmt.Print("Player ", current_turn, " to play:")
		//_, err := fmt.Scan(&player_input)
		player_input, err = reader.ReadString('\n')
		if err != nil {
			return
		}

		indexes := strings.Split(strings.TrimSpace(player_input), " ")

		x, _ := strconv.Atoi(indexes[0])
		y, _ := strconv.Atoi(indexes[1])

		if !is_legal(board, x, y) {
			fmt.Println("Last move", strings.TrimSpace(player_input), "by", current_turn, "was illegal")
			if current_turn == "X" {
				fmt.Println("Player 0 won")
			} else {
				fmt.Println("Player X won")
			}
			break
		}

		board[x][y] = current_turn
		winner := get_winner_or_nil(board)
		if winner != nil {
			fmt.Printf("Player %s won\n", winner)
			break
		} else {
			if turns == 9 {
				fmt.Println("No winner")
			}
		}

		if current_turn == p1 {
			current_turn = p2
		} else {
			current_turn = p1
		}
	}

	print_board(board)
	fmt.Println("Game Finished")
}
