package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// RetabConfig instructions for @Retab
type RetabConfig struct {
	// FromIndent defines how many spaces an indentation level took in the file had before
	FromIndent int
	// FromIndent defines how many spaces an indentation level will take in the file after
	ToIndent int
	// TabSpaceCount defines how many spaces a tab character is worth
	TabSpaceCount int
}

// Retab changes files from a given indentation size/char to a different one
func Retab(inputFilePath string, cfg RetabConfig) error {
	var err error

	fromStdin := inputFilePath == "-"

	var fif *os.File
	if !fromStdin {
		// resolve filePath
		inputFilePath, err = filepath.Abs(inputFilePath)
		if err != nil {
			return err
		}

		fif, err = os.Open(inputFilePath)
		if err != nil {
			return err
		}
	}

	var result bytes.Buffer
	{ // convert indentation and write to tempfile
		var scanner *bufio.Scanner

		if fromStdin {
			scanner = bufio.NewScanner(os.Stdin)
		} else {
			scanner = bufio.NewScanner(fif)
		}

		defer func() {
			if !fromStdin {
				fif.Close()
			}
		}()

		for scanner.Scan() {
			line := scanner.Text()

			if line == "" { // skip empty lines
				result.WriteString("\n")
				continue
			}

			indentCount := 0

			firstNonIdentIndex := 0
			c := ' '
		countIndentationBySpaces:
			for firstNonIdentIndex, c = range line {
				switch c {
				case ' ':
					indentCount++
				case '\t':
					indentCount += cfg.TabSpaceCount
				default:
					break countIndentationBySpaces
				}
			}

			newIndent := 0
			oldIndent := indentCount
			for oldIndent > 0 {
				if oldIndent >= cfg.FromIndent {
					oldIndent -= cfg.FromIndent
					newIndent += cfg.ToIndent
				} else {
					newIndent += oldIndent
					oldIndent = 0
				}
			}

			result.WriteString(strings.Repeat(" ", newIndent))
			result.WriteString(line[firstNonIdentIndex:])
			result.WriteString("\n")
		}
		if err := scanner.Err(); err != nil {
			return err
		}
	}

	if fromStdin {
		os.Stdout.Write(result.Bytes())
	} else {
		err = ioutil.WriteFile(inputFilePath, result.Bytes(), 0755)
		if err != nil {
			return err
		}
	}

	return nil
}
