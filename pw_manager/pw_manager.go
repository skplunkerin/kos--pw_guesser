package pw_manager

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// RemoveFirstLineFromFile copies the first non-empty line from the file,
// removes it (and any blank lines above it) from the file, then returns the
// line or err.
func RemoveFirstLineFromFile(fromFilePath string) (string, error) {
	// Extract the file name from the path
	fromFileName := filepath.Base(fromFilePath)

	// Read the contents of file
	file, err := os.Open(fromFilePath)
	if err != nil {
		fmt.Printf("Error opening file %s: %s", fromFilePath, err.Error())
		return "", err
	}
	defer file.Close()

	// TODO: update this to not read the entire file into memory. [skplunkerin]
	fileContent := make([]byte, 0)
	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err != nil {
			break
		}
		fileContent = append(fileContent, buf[:n]...)
	}

	// Split the content into lines
	lines := strings.Split(string(fileContent), "\n")

	// Find the index of the first non-empty line
	firstNonEmptyIndex := -1
	for i, line := range lines {
		if strings.TrimSpace(line) != "" {
			firstNonEmptyIndex = i
			break
		}
	}

	if firstNonEmptyIndex == -1 {
		fmt.Printf("No non-empty lines found in file %s.", fromFileName)
		return "", nil
	}

	// Create a new slice of lines excluding the first non-empty line and preceding blank lines
	remainingLines := lines[firstNonEmptyIndex+1:]

	// Write the first line to file2.txt
	firstLine := strings.TrimSpace(lines[firstNonEmptyIndex])
	if firstLine == "" {
		fmt.Printf("No non-empty lines found in file %s.", fromFileName)
		return "", nil
	}

	// Write back the remaining lines to file
	remainingContent := strings.Join(remainingLines, "\n")
	err = os.WriteFile(fromFilePath, []byte(remainingContent), 0644)
	if err != nil {
		fmt.Printf("Error writing to file %s: %s", fromFilePath, err.Error())
		return "", err
	}
	return firstLine, nil
}

// AppendLineToFile saves the line to the end of the file.
func AppendLineToFile(line, toFilePath string) error {
	// create destination file if it doesn't exist, else open it
	var (
		file *os.File
		err  error
	)
	_, err = os.Stat(toFilePath)
	if os.IsNotExist(err) {
		file, err = os.Create(toFilePath)
		if err != nil {
			fmt.Printf("Error creating file %s: %s", toFilePath, err.Error())
			return err
		}
	} else {
		file, err = os.OpenFile(toFilePath, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening file %s: %s", toFilePath, err.Error())
			return err
		}
	}
	defer file.Close()
	// Prepend the line to file
	if line == "" {
		// fmt.Println("empty line found, skipping.")
		return nil
	}
	_, err = file.WriteString(line + "\n")
	if err != nil {
		fmt.Printf("Error prepending to file %s: %s", toFilePath, err.Error())
		return err
	}
	// Extract the file name from the path
	// toFileName := filepath.Base(toFilePath)
	// fmt.Printf("Line prepended to file %s:", toFileName)
	return nil
}
