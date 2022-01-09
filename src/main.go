// Golang program to show how
// to use command-line arguments
package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Get environment variable newmove
	newmoveString := os.Getenv("ISSUE_TITLE")
	winnerName := os.Getenv("ISSUE_CREATOR")
	newmove, error := strconv.Atoi(newmoveString)
	if error != nil {
		os.Stdout.WriteString("Error: Invalid move")
		panic("Invalid move")
	}
	//read README.md file
	readmeFile, error := ioutil.ReadFile("README.md")
	if error != nil {
		panic("Error: Unable to open README.md file")
	}
	readme := string(readmeFile)

	// Get text inside <!-- START: tic-tac-toe --> and <!-- END: tic-tac-toe --> comment
	board := strings.Split(strings.Split(readme, "<!-- START: tic-tac-toe -->")[1], "<!-- END: tic-tac-toe -->")[0]

	//Revmoe useless table indentation
	board = strings.ReplaceAll(board, "|-|-|-|", "")
	board = strings.ReplaceAll(board, "\n", "")
	board = strings.ReplaceAll(board, "	", "")

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
	if xCount <= oCount {
		XorO = "[![X](https://raw.githubusercontent.com/Lettly/Lettly/main/media/x.svg)](#)"
	} else {
		XorO = "[![O](https://raw.githubusercontent.com/Lettly/Lettly/main/media/o.svg)](#)"
	}

	// Check if newmove is valid
	if newmove >= 0 && newmove <= 9 {
		if strings.Contains(boardArray[newmove], "[O]") || strings.Contains(boardArray[newmove], "[X]") {
			os.Stdout.WriteString("Error: Invalid move")
			panic("Invalid move")
		}
		boardArray[newmove] = XorO
	} else {
		os.Stdout.WriteString("Error: Invalid move")
		panic("Invalid move")
	}

	//Check if game is over
	winner := ""
	if boardArray[0] == boardArray[1] && boardArray[1] == boardArray[2] {
		winner = boardArray[0]
	} else if boardArray[3] == boardArray[4] && boardArray[4] == boardArray[5] {
		winner = boardArray[3]
	} else if boardArray[6] == boardArray[7] && boardArray[7] == boardArray[8] {
		winner = boardArray[6]
	} else if boardArray[0] == boardArray[3] && boardArray[3] == boardArray[6] {
		winner = boardArray[0]
	} else if boardArray[1] == boardArray[4] && boardArray[4] == boardArray[7] {
		winner = boardArray[1]
	} else if boardArray[2] == boardArray[5] && boardArray[5] == boardArray[8] {
		winner = boardArray[2]
	} else if boardArray[0] == boardArray[4] && boardArray[4] == boardArray[8] {
		winner = boardArray[0]
	} else if boardArray[2] == boardArray[4] && boardArray[4] == boardArray[6] {
		winner = boardArray[2]
	} else if xCount+oCount == 8 {
		winner = "draw"
	}
	if strings.Contains(winner, "[X]") {
		winner = "X"
	} else if strings.Contains(winner, "[O]") {
		winner = "O"
	}

	// if game is over, reset board
	if winner != "" {
		for i := 0; i < 9; i++ {
			boardArray[i] = "[![X clickable](https://raw.githubusercontent.com/Lettly/Lettly/main/media/x-trasparent.svg)](https://github.com/Lettly/Lettly/issues/new?body=Please+do+not+change+the+title.+Just+click+%22Submit+new+issue%22.+You+don%27t+need+to+do+anything+else+%3AD&title=" + strconv.Itoa(i) + ")"
		}
	} else {
		// Change all the gray boxes to the opposite sign (X or O)
		for i := 0; i < 9; i++ {
			if strings.Contains(boardArray[i], "[X clickable]") {
				boardArray[i] = "[![O clickable](https://raw.githubusercontent.com/Lettly/Lettly/main/media/o-trasparent.svg)](https://github.com/Lettly/Lettly/issues/new?body=Please+do+not+change+the+title.+Just+click+%22Submit+new+issue%22.+You+don%27t+need+to+do+anything+else+%3AD&title=" + strconv.Itoa(i) + ")"
			} else if strings.Contains(boardArray[i], "[O clickable]") {
				boardArray[i] = "[![X clickable](https://raw.githubusercontent.com/Lettly/Lettly/main/media/x-trasparent.svg)](https://github.com/Lettly/Lettly/issues/new?body=Please+do+not+change+the+title.+Just+click+%22Submit+new+issue%22.+You+don%27t+need+to+do+anything+else+%3AD&title=" + strconv.Itoa(i) + ")"
			}
		}
	}

	//Prepare board for output
	output := strings.Split(readme, "<!-- START: tic-tac-toe -->")[0]
	output += "<!-- START: tic-tac-toe -->\n"
	for i := 0; i < 9; i++ {
		output += boardArray[i] + "|"
		if i == 2 {
			output += "\n|-|-|-|\n"
		}
		if i == 5 || i == 8 {
			output += "\n"
		}
	}
	output += "<!-- END: tic-tac-toe -->"
	if winner == "X" || winner == "O" {
		output += `<!-- START: tic-tac-toe-winner -->Last game was won by <b>` + winnerName + ` (` + winner + `)</b><!-- END: tic-tac-toe-winner -->`
		output += strings.Split(readme, "<!-- END: tic-tac-toe-winner -->")[1]
	} else if winner == "draw" {
		output += `<!-- START: tic-tac-toe-winner -->Last game was a draw<!-- END: tic-tac-toe-winner -->`
		output += strings.Split(readme, "<!-- END: tic-tac-toe-winner -->")[1]
	} else {
		output += strings.Split(readme, "<!-- END: tic-tac-toe -->")[1]
	}

	os.Stdout.WriteString(output)
	// write to file
	err := ioutil.WriteFile("README.md", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}
