package main

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/atotto/clipboard"
	"github.com/skplunkerin/kos--pw_guesser/pw_manager"
)

var (
	// Testing:
	// srcFileName = "testFile1.txt"
	// toFileName  = "testFile2.txt"
	srcFileName  = "1combinations.txt"
	toFileName   = "2attempts.txt"
	notification = Notification{
		display: false,
		delay:   1 * time.Second,
		Title:   "Notification from Go",
		Message: "message not set...",
	}
)

type Notification struct {
	display bool
	delay   time.Duration
	Title   string
	Message string
}

func main() {
	// TODO:
	//  - add logic to call this if the `1combinations.txt` file doesn't exist or
	//    is empty.
	//  - limit combinations to including each number 1+ times.
	// createCombinations([]int{3, 4, 5, 7}, 5)

	line, err := pw_manager.RemoveFirstLineFromFile(srcFileName)
	if err != nil {
		fmt.Printf("CopyFirstLine() failed: %s", err.Error())
		return
	}

	copyToClipboard(line)

	err = pw_manager.PrependLineToFile(line, toFileName)
	if err != nil {
		fmt.Printf("SaveLineToFile() failed: %s", err.Error())
		return
	}
}

// generateCombinations generates all combinations of numbers of a given length
// from a given set of numbers.
func createCombinations(numbers []int, length int) {
	// generate combinations of numbers and save them to a file
	var combinations [][]int
	generateCombinations(numbers, length, nil, &combinations)

	err := saveCombinationsToFile(srcFileName, "", combinations)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Combinations saved to %s", srcFileName)
}

func (n Notification) displayNotification() {
	if !n.display {
		return
	}
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`display notification "%s" with title "%s"`, n.Message, n.Title))
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error displaying notification:", err)
	}
}

// copyToClipboard copies a string to the clipboard.
func copyToClipboard(str string) {
	err := clipboard.WriteAll(str)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		if notification.display {
			notification.Message = fmt.Sprintf("'%s' copied", str)
			notification.displayNotification()
			// Give some time for the notification to be displayed before the program
			// exits
			time.Sleep(notification.delay)
		}
	}
}
