package main

func GetBestMove(board [3][3]int, difficulty int) (int, int) {
	// if difficulty == 1 {
	// 	return getRandomMove(board)
	// } else if difficulty == 2 {
	// 	return getMediumMove(board)
	// }
	return minimaxMove(board)
}

func minimaxMove(board [3][3]int) (int, int) {
	bestScore := -1000
	var move [2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == Empty {
				board[i][j] = AI
				score := minimax(board, 0, false)
				board[i][j] = Empty
				if score > bestScore {
					bestScore = score
					move = [2]int{i, j}
				}
			}
		}
	}
	return move[0], move[1]
}

func minimax(board [3][3]int, depth int, isMax bool) int {
	winner := CheckWinner(board)
	if winner != 0 {
		return winner * -1
	}
	if IsBoardFull(board) {
		return 0
	}

	if isMax {
		best := -1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = AI
					best = max(best, minimax(board, depth+1, false))
					board[i][j] = Empty
				}
			}
		}
		return best
	} else {
		best := 1000
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if board[i][j] == Empty {
					board[i][j] = Player
					best = min(best, minimax(board, depth+1, true))
					board[i][j] = Empty
				}
			}
		}
		return best
	}
}

func CheckWinner(board [3][3]int) int {
	for i := 0; i < 3; i++ {
		if board[i][0] != Empty && board[i][0] == board[i][1] && board[i][1] == board[i][2] {
			return board[i][0]
		}
		if board[0][i] != Empty && board[0][i] == board[1][i] && board[1][i] == board[2][i] {
			return board[0][i]
		}
	}
	if board[0][0] != Empty && board[0][0] == board[1][1] && board[1][1] == board[2][2] {
		return board[0][0]
	}
	if board[0][2] != Empty && board[0][2] == board[1][1] && board[1][1] == board[2][0] {
		return board[0][2]
	}
	return 0
}

func IsBoardFull(board [3][3]int) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == Empty {
				return false
			}
		}
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
