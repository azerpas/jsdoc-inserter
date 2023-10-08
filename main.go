package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <filename.js> <line_to_add>")
		os.Exit(1)
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	lineToAdd := os.Args[2]

	scanner := bufio.NewScanner(file)
	var lines []string
	var inComment bool
	var commentLines []string

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)

		if strings.HasPrefix(strings.TrimSpace(line), "/**") {
			inComment = true
			commentLines = nil
		}

		if inComment {
			commentLines = append(commentLines, line)
		}

		if strings.HasPrefix(strings.TrimSpace(line), "*/") {
			inComment = false
			if !lineExistsInComment(commentLines, lineToAdd) {
				lines = appendCommentLine(lines, commentLines, lineToAdd)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	// Open file for writing
	fileName := strings.Replace(os.Args[1], ".js", "_new.js", 1)
	outputFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Error writing to file:", err)
			os.Exit(1)
		}
	}
	writer.Flush()
}

func lineExistsInComment(commentLines []string, lineToAdd string) bool {
	for _, line := range commentLines {
		if strings.Contains(line, lineToAdd) {
			return true
		}
	}
	return false
}

func appendCommentLine(lines []string, commentLines []string, lineToAdd string) []string {
	indentation := getIndentation(commentLines[0])
	newLine := indentation + " * " + lineToAdd

	// Find the index of the comment block's end line within the lines slice
	endIndex := len(lines) - 1
	for i := endIndex; i >= 0; i-- {
		if strings.TrimSpace(lines[i]) == "*/" {
			endIndex = i
			break
		}
	}

	// Insert the new line before the comment block's end line
	newLines := append(lines[:endIndex], append([]string{newLine}, lines[endIndex:]...)...)
	return newLines
}

func getIndentation(line string) string {
	return regexp.MustCompile(`^(\s*)`).FindString(line)
}
