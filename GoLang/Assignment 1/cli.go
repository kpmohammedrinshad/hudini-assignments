// Objective:Create a command-line interface (CLI) application in Go that allows users to input a block of text and
// then calculates the frequency of each word in the text. This project will help you understand
// and implement basic Go concepts like variables, loops, conditionals, maps, functions, and strings.
// Requirements:
// Input Text: Allow users to input a block of text.
// Process Text: Count the frequency of each word in the text.
// Display Frequencies: Display the word frequencies in a readable format.
// Exit: Allow the user to exit the application.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func FrequencyWord() {
	// take input text from user
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter text:")
	text, _ := reader.ReadString('\n')
	//store the text to the map
	freqWord := map[string]int{}
	//split the text into words
	text = strings.TrimSpace(text)
	words := strings.Split(text, " ")
	// words:=strings.Fields(text)
	//counting the frequency
	for _, word := range words {
		freqWord[word]++
	}
	//print the map
	for w, f := range freqWord {
		fmt.Printf("'%s'- %d\n", w, f)
	}

}

func main() {
	FrequencyWord()
}
