package main

import (
	"os"
	"bufio"
	"fmt"
	"strings"
	"strconv"
)

// GameBoard is a two-dimensional slice of booleans.
type GameBoard [][]bool

// InField returns true if the cell at (i, j) is within the boundaries of board.
func InField(board GameBoard, i, j int) bool {
	var rows int
	var columns int

	// Count the rows.
	for x := 0; x < len(board); x++ {
		rows++
	}
	// Count the columns.
	for y := 0; y < len(board[0]); y++ {
		columns++
	}
	if i >= 0 && i < rows && j >= 0 && j < columns {
		return true
	}
	return false
}

// CountRows returns the number of rows in board.
func CountRows(board GameBoard) int {
	return len(board)
}

// CountCols returns the number of columns in board (assumes rectangular board).
func CountCols(board GameBoard) int {
	if CountRows(board) == 0 {
		panic("Error: empty board given to CountCols")
	}
	return len(board[0])
}

// CountLiveNeighbors counts the live neighbors of the cell at (r, c) in board.
func CountLiveNeighbors(board GameBoard, r, c int) int {
	lifecells := 0
	for i := r - 1; i <= r+1; i++ {
		for j := c - 1; j <= c+1; j++ {
			if InField(board, i, j) {
				if board[i][j] {
					lifecells++
				}
			}
		}
	}
	// Subtract the cell itself if it is alive.
	if board[r][c] {
		lifecells--
	}
	return lifecells
}

// UpdateCell returns the updated state (true or false) for the cell at (r, c)
// based on the Game of Life rules.
func UpdateCell(board GameBoard, r, c int) bool {
	var status bool
	if InField(board, r, c) {
		lives := CountLiveNeighbors(board, r, c)
		if board[r][c] {
			if lives == 2 || lives == 3 {
				status = true
			} else {
				status = false
			}
		} else {
			if lives == 3 {
				status = true
			} else {
				status = false
			}
		}
	}
	return status
}

// CopyBoard creates and returns a deep copy of the provided board.
func CopyBoard(original GameBoard) GameBoard {
	rows := CountRows(original)
	cols := CountCols(original)

	newBoard := InitializeBoard(rows, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			newBoard[i][j] = original[i][j]
		}
	}
	return newBoard
}

// UpdateBoard returns a new board after updating each cell in currBoard.
func UpdateBoard(currBoard GameBoard) GameBoard {
	grid := CopyBoard(currBoard)
	rows := CountRows(currBoard)
	cols := CountCols(currBoard)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			grid[i][j] = UpdateCell(currBoard, i, j)
		}
	}
	return grid
}

// PlayGameOfLife runs the Game of Life for numGens generations starting from initialBoard
// and returns a slice of boards (one for each generation, including the first update).
func PlayGameOfLife(initialBoard GameBoard, numGens int) []GameBoard {
	boards := make([]GameBoard, numGens+1)
	current := initialBoard
	boards[0] = UpdateBoard(initialBoard)
	for i := 1; i <= numGens; i++ {
		boards[i] = UpdateBoard(current)
		current = boards[i]
	}
	return boards
}

// InitializeBoard creates a new game board with the given number of rows and columns,
// with all cells initialized to false.
func InitializeBoard(numRows, numCols int) GameBoard {
	board := make(GameBoard, numRows)
	for r := range board {
		board[r] = make([]bool, numCols)
	}
	return board
}
