package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

var pl = fmt.Println

// // Hangman game
/*
	    +---+
	    0	  |
	   /|\  |
	   / \  |
	   	  ===

	Secret Word : M_____
	Incorrect Guesses : A

	Guess a Letter : Y

	Sorry you're Dead! The word is MONKEY
	Yes the Secret Word is MONKEY

	Please Enter Only One Letter
	Please Enter a Letter
	Please Enter a Letter you Haven't Guessed
*/

var hmArr = [7]string{
	" +---+\n" +
		"     |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		"	0   |\n" +
		"     |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		"	0   |\n" +
		" |   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		"	0   |\n" +
		"/|   |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		"	0   |\n" +
		"/|\\ |\n" +
		"     |\n" +
		"    ===\n",
	" +---+\n" +
		"	0   |\n" +
		"/|\\ |\n" +
		"/    |\n" +
		"    ===\n",
	" +---+\n" +
		"	0   |\n" +
		"/|\\ |\n" +
		"/ \\ |\n" +
		"    ===\n",
}

var wordArr = [7]string{"JAZZ", "ZIGZAG", "ZILCH", "ZIPPER", "ZODIAC", "ZOMBIE", "FLUFF"}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

var alphabet = "aAbBcCçÇdDeEfFgGhHiIjJkKlLmMnNoOpPqQrRsStTuUvVwWxXyYzZ"

func main() {
	randWord = wordArr[0]

	for true {
		// Show Game Board
		showGameBoard()

		// Get a letter from the user
		pl("Guess a letter:")
		reader := bufio.NewReader(os.Stdin)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess, err = validateGuess(guess)
		if err != nil {
			pl(err)
			continue
		}

		// A. If the letter is in the word
		if strings.Contains(randWord, guess) {
			correctLetters = append(correctLetters, guess)
			guessedLetters += guess
			pl("Correct Letter")
		}
		// 1. Are there more letters to guess?
		// 2. If no more letters to guess (YOU WIN)
		if checkVictory() {
			pl("Yes the Secret Word is " + randWord)
			return
		}
		// B. If the letter is NOT in the word
		// 1. Add new letter to guessedLetters
		// wrongGuesses
		// 2. If no more chances (YOU LOSE)
	}

}

func showGameBoard() {
	pl(hmArr[len(wrongGuesses)])
	fmt.Print("Secret Word :")
	for _, letter := range randWord {
		if strings.Contains(guessedLetters, string(letter)) {
			fmt.Print(string(letter))
		} else {
			fmt.Print("_")
		}
	}
	pl()
}

func validateGuess(guess string) (letter string, err error) {
	guess = strings.TrimSpace(guess)

	isValidGuess := strings.Contains(alphabet, guess)

	if len(guess) != 1 {
		return "", errors.New("Please enter only one letter")
	}

	if !isValidGuess {
		return "", errors.New("You should enter a letter")
	}

	return guess, nil
}

func checkVictory() bool {
	for _, letter := range randWord {
		if !strings.Contains(guessedLetters, string(letter)) {
			return false
		}
	}
	return true
}
