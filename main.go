/*
 *Read input from STDIN. Print your output to STDOUT
 *Use fmt.Scanf to read input from STDIN and fmt. Println to write output to STDOUT
 */

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// assuming input comes from console
	var numOfPkts int
	var numOfBallsText string
	// scan input
	fmt.Scan(&numOfPkts)
	fmt.Scanf("%s", &numOfBallsText)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numOfBallsText = scanner.Text()
	// format input: space based input
	splittedText := strings.Split(numOfBallsText, " ") //
	var numOfBalls []int
	// convert string to int
	for _, eachText := range splittedText {
		input, err := strconv.Atoi(eachText)
		if err == nil {
			numOfBalls = append(numOfBalls, input)
		}
	}
	moves := 0
	for _, eachNumOfBalls := range numOfBalls {
		moves = moves + divideToGroup(eachNumOfBalls)
	}
	fmt.Println("Maximum Moves:", moves)
}

func validateInputs(numOfPkts int, numOfBalls []int) error {
	var err error
	// validate pkt numbers
	if numOfPkts < 0 || numOfPkts > 100 {
		return errors.New("Invalid number of packets")
	}
	// validate balls
	for _, eachNumOfBalls := range numOfBalls {
		if eachNumOfBalls < 0 || eachNumOfBalls > (10*10) {
			err = errors.New("invalid number of balls")
			break
		}
	}
	return err
}

func divideToGroup(numOfBalls int) int {
	var secondMove, thirdMove int
	if numOfBalls == 1 {
		return 1 // one doesnt require any processing
	}
	// find the least divisibility factor
	var factor int
	// for even number it will be 2
	if numOfBalls%2 == 0 {
		factor = 2
	} else {
		// for odd number it has determined from 3 to ++2 increment
		divisibility := 3
		for {
			if numOfBalls%divisibility == 0 {
				factor = divisibility
				break
			}
			divisibility = divisibility + 2
		}
	}
	// determine the number of moves
	secondMove = numOfBalls / factor
	thirdMove = numOfBalls
	return 1 + secondMove + thirdMove

}
