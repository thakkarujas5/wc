// package main

// import (
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	argsLength := len(os.Args)

// 	var filesLength int
// 	lineCount := false
// 	charCount := false
// 	wordCount := false

// 	for x, i := range os.Args {

// 		if i == "-l" {
// 			lineCount = true
// 		}

// 		if i == "-c" {
// 			charCount = true
// 		}
// 		if i == "-w" {
// 			wordCount = true
// 		}
// 		if i != "./wc" && i != "-l" && i != "-w" && i != "-c" {
// 			filesLength = argsLength - x
// 			break
// 		}
// 	}

// 	files := make([]string, filesLength)

// 	for y, x := 0, argsLength-filesLength; x < argsLength; x, y = x+1, y+1 {
// 		files[y] = os.Args[x]
// 	}

// 	for _, file := range files {

// 		content, err := os.ReadFile(file)

// 		if err != nil {
// 			fmt.Println(err)
// 			return
// 		}

// 		fileContent := string(content)

// 		lines := strings.Split(fileContent, "\n")

// 		wc := 0
// 		cc := 0
// 		for _, line := range lines {

// 			words := strings.Split(line, " ")
// 			wc += len(words)

// 			for _, word := range words {
// 				cc += len(word)
// 			}
// 		}

// 		if lineCount {
// 			fmt.Print(len(lines), " ")
// 		}

// 		if wordCount {
// 			fmt.Print(wc, " ")
// 		}

// 		if charCount {
// 			fmt.Print(cc, " ")
// 		}

// 		fmt.Println(file)
// 	}

// }
