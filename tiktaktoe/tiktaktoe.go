package main

import (
	"fmt"
	"math/rand"
)

type Board [9]string

func newBoard() Board {
	var b Board
	for i := range b {
		b[i] = fmt.Sprintf("%d", i+1)
	}
	return b
}

func (b *Board) print() {
	fmt.Printf("\n %s | %s | %s\n", b[0], b[1], b[2])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s\n", b[3], b[4], b[5])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s\n\n", b[6], b[7], b[8])
}

var winLines = [8][3]int{
	{0, 1, 2}, {3, 4, 5}, {6, 7, 8}, // rows
	{0, 3, 6}, {1, 4, 7}, {2, 5, 8}, // cols
	{0, 4, 8}, {2, 4, 6}, // diagonals
}

func (b *Board) winner() string {
	for _, line := range winLines {
		a, c, d := b[line[0]], b[line[1]], b[line[2]]
		if a == c && c == d {
			return a
		}
	}
	return ""
}

func (b *Board) isFull() bool {
	for _, cell := range b {
		if cell != "X" && cell != "O" {
			return false
		}
	}
	return true
}

func (b *Board) availableMoves() []int {
	var moves []int
	for i, cell := range b {
		if cell != "X" && cell != "O" {
			moves = append(moves, i)
		}
	}
	return moves
}

func (b *Board) playerMove() {
	for {
		fmt.Print("Your move (1-9): ")
		var input int
		_, err := fmt.Scan(&input)
		if err != nil || input < 1 || input > 9 {
			fmt.Println("Invalid input. Enter a number 1-9.")
			continue
		}
		idx := input - 1
		if b[idx] == "X" || b[idx] == "O" {
			fmt.Println("That cell is taken. Try again.")
			continue
		}
		b[idx] = "X"
		return
	}
}

func (b *Board) computerMove() {
	// Try to win
	if idx := b.findBestMove("O"); idx >= 0 {
		b[idx] = "O"
		return
	}
	// Block player
	if idx := b.findBestMove("X"); idx >= 0 {
		b[idx] = "O"
		return
	}
	// Take center
	if b[4] != "X" && b[4] != "O" {
		b[4] = "O"
		return
	}
	// Random move
	moves := b.availableMoves()
	b[moves[rand.Intn(len(moves))]] = "O"
}

func (b *Board) findBestMove(mark string) int {
	for _, line := range winLines {
		a, c, d := b[line[0]], b[line[1]], b[line[2]]
		cells := []string{a, c, d}
		count, empty := 0, -1
		for j, v := range cells {
			if v == mark {
				count++
			} else if v != "X" && v != "O" {
				empty = line[j]
			}
		}
		if count == 2 && empty >= 0 {
			return empty
		}
	}
	return -1
}

func main() {
	fmt.Println("Tic-Tac-Toe: You are X, Computer is O")
	fmt.Println("Enter the number of the cell you want to play.")

	board := newBoard()
	board.print()

	for {
		board.playerMove()
		board.print()
		if w := board.winner(); w != "" {
			fmt.Println("You win!")
			return
		}
		if board.isFull() {
			fmt.Println("It's a draw!")
			return
		}

		fmt.Println("Computer's turn...")
		board.computerMove()
		board.print()
		if w := board.winner(); w != "" {
			fmt.Println("Computer wins!")
			return
		}
		if board.isFull() {
			fmt.Println("It's a draw!")
			return
		}
	}
}
