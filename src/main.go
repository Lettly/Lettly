// Golang program to show how
// to use command-line arguments
package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	//Get environment variable newmove
	newmoveString := os.Getenv("newmove")
	newmove, error := strconv.Atoi(newmoveString)
	if error != nil {
		os.Stdout.WriteString("Error: Invalid move")
		os.Exit(1)
	}
	//Get environment variable readme
	readme := os.Getenv("readme")
	// Get board
	// Get text inside <!-- START: tick-tack-toe --> and <!-- END: tick-tack-toe --> comment
	board := strings.Split(strings.Split(readme, "<!-- START: tick-tack-toe -->")[1], "<!-- END: tick-tack-toe -->")[0]

	//Revmoe useless table indentation
	board = strings.ReplaceAll(board, "|-|-|-|", "")

	// Convert board to array
	boardArray := strings.Split(board, "|")

	// Get if next move is X or O
	xCount := 0
	oCount := 0
	XorO := ""
	for i := 0; i < 9; i++ {
		if strings.Contains(boardArray[i], "[X]") {
			xCount++
		} else if strings.Contains(boardArray[i], "[O]") {
			oCount++
		}
	}
	if xCount >= oCount {
		XorO = "![X](https://raw.githubusercontent.com/Lettly/Lettly/main/media/x.svg)"
	} else {
		XorO = "![O](https://raw.githubusercontent.com/Lettly/Lettly/main/media/o.svg)"
	}

	// Check if newmove is valid
	if newmove >= 1 && newmove <= 9 {
		if strings.Contains(boardArray[newmove-1], "[O]") || strings.Contains(boardArray[newmove-1], "[X]") {
			os.Stdout.WriteString("Error: Invalid move")
			os.Exit(1)
		}
		boardArray[newmove-1] = XorO
	} else {
		os.Stdout.WriteString("Error: Invalid move")
		os.Exit(1)
	}
}
