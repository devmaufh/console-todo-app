package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// sanitizeConsoleInput remove the line breaks of the given string
func sanitizeConsoleInput(s string) string {
	s = strings.Replace(s, "\r\n", "", -1)
	s = strings.Replace(s, "\n", "", -1)

	return s
}

// ReadConsoleInput read console input and returns user input
func ReadConsoleInput(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(message)
	fmt.Print("-> ")
	input, _ := reader.ReadString('\n')
	return sanitizeConsoleInput(input)
}
