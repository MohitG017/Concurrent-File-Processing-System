package task

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	FilePath string
	Result   string
}

// ProcessFile processes the file and counts occurrences of the word
func (t *Task) ProcessFile(word string) (int, error) {
	file, err := os.Open(t.FilePath)
	if err != nil {
		return 0, fmt.Errorf("error opening file %s: %v", t.FilePath, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0
	for scanner.Scan() {
		line := scanner.Text()
		count += strings.Count(line, word)
	}

	if err := scanner.Err(); err != nil {
		return 0, fmt.Errorf("error reading file %s: %v", t.FilePath, err)
	}

	return count, nil
}
