package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

type flagOptions struct {
	lineFlag bool
	wordFlag bool
	charFlag bool
}

type result struct {
	lineCount int
	wordCount int
	charCount int
	filename  string
	err       error
}

var (
	flagSet                                        flagOptions
	totalLineCount, totalWordCount, totalCharCount int
)

const maxOpenFileLimit = 10

func init() {
	// Add flags to count lines, words, and characters
	rootCmd.Flags().BoolVarP(&flagSet.lineFlag, "lines", "l", false, "Count number of lines")
	rootCmd.Flags().BoolVarP(&flagSet.wordFlag, "words", "w", false, "Count number of words")
	rootCmd.Flags().BoolVarP(&flagSet.charFlag, "chars", "c", false, "Count number of characters")
}

var rootCmd = &cobra.Command{
	Use:   "wc",
	Short: "A brief description of your application",
	Long:  `A longer description that spans multiple lines and likely contains examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var wg sync.WaitGroup
		maxOpenFilesLimitBuffer := make(chan int, maxOpenFileLimit)

		for _, file := range args {
			go worker(file, &wg, maxOpenFilesLimitBuffer)
			wg.Add(1)
		}

		wg.Wait()
	},
}

func worker(file string, wg *sync.WaitGroup, maxOpenFilesLimitBuffer chan int) {

	lines := make(chan string)
	errChan := make(chan error)

	maxOpenFilesLimitBuffer <- 1
	defer func() {
		wg.Done()
		<-maxOpenFilesLimitBuffer
	}()

	go readLinesInFile(file, lines, errChan)

	result := count(lines, errChan)
}

func readLinesInFile(filename string, lines chan<- string, errChan chan<- error) {

	var scanner *bufio.Scanner

	const chunkSize = 1024 * 1024

	defer close(lines)
	defer close(errChan)

	file, err := os.Open(filename)

	if err != nil {

		errChan <- err
	}

	defer file.Close()

	scanner = bufio.NewScanner(file)
	scanner.Buffer(make([]byte, chunkSize), chunkSize)

	for scanner.Scan() {
		lines <- scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		err = fmt.Errorf(
			"gowc: " + strings.Replace(err.Error(), "read ", "", 1) + "\n",
		)
		errChan <- err
	}

}

func count(lines <-chan string, errChan <-chan error) result {
	var r result

	for {
		select {
		case err := <-errChan:
			if err != nil {
				r.err = err
				errChan = nil
				return r
			}
		}
	case line, ok := <-lines:
		if !ok {
			return r
		}
	}

}

func main() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
