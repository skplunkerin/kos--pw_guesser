package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/atotto/clipboard"
	"github.com/skplunkerin/kos--pw_guesser/pw_manager"
)

// FilesInfo contains the paths to the files to be used.
type FilesInfo struct {
	Path        string // the relative path from the current working directory
	SrcFileName string
	ToFileName  string
}

var (
	dirt = "TOPSECRT"

	filePaths = map[string]FilesInfo{
		// not an actual dirt, just for testing purposes
		"testing": {
			Path:        "/",
			SrcFileName: "testFile1.txt",
			ToFileName:  "testFile2.txt",
		},
		"phlosphy": {
			Path:        "dirts/phlosphy/combinations/",
			SrcFileName: "1combinations.txt",
			ToFileName:  "2attempts.txt",
		},
		"TOPSECRT": {
			Path:        "dirts/TOPSECRT/combinations/",
			SrcFileName: "1combinations.txt",
			ToFileName:  "2attempts.txt",
		},
	}
)

func main() {
	// Get the current working directory
	currentWorkingDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting working directory:", err)
		return
	}
	srcFilePath := filepath.Join(currentWorkingDir, filePaths[dirt].Path, filePaths[dirt].SrcFileName)
	toFilePath := filepath.Join(currentWorkingDir, filePaths[dirt].Path, filePaths[dirt].ToFileName)
	// TODO: only print this out during `go build` and `go run` so users know the
	// files being used in final binary. [skplunkerin]
	var (
		building = false
		running  = false
	)
	if building || running {
		fmt.Printf("currentWorkingDir: %s\n", currentWorkingDir)
		fmt.Printf("srcFilePath: %s\n", srcFilePath)
		fmt.Printf("toFilePath: %s\n", toFilePath)
	}

	if dirt == "TOPSECRT" {
		// TODO: only print this out during `go build` and `go run`. [skplunkerin]
		if building || running {
			fmt.Printf("Checking if combinations file generation is needed for dirt %s...\n", dirt)
		}
		// Check if the file exists or is empty
		srcFileInfo, err := os.Stat(srcFilePath)
		if os.IsNotExist(err) || srcFileInfo.Size() == 0 {
			if building || running {
				fmt.Printf("Combinations file needed, generating...\n")
			}
			createIntCombinations(srcFilePath, []int{3, 4, 5, 7}, 5)
		} else if err != nil {
			fileName := filepath.Base(srcFilePath)
			fmt.Printf("Error checking file name %s: %s\n", fileName, err.Error())
		} else {
			if building || running {
				fmt.Printf("Combinations file exists and is not empty, skipping generation\n")
			}
		}
	}

	line, err := pw_manager.RemoveFirstLineFromFile(srcFilePath)
	if err != nil {
		fmt.Printf("CopyFirstLine() failed: %s", err.Error())
		return
	}
	copyToClipboard(line)
	err = pw_manager.AppendLineToFile(line, toFilePath)
	if err != nil {
		fmt.Printf("SaveLineToFile() failed: %s", err.Error())
		return
	}
}

// copyToClipboard copies a string to the clipboard.
func copyToClipboard(str string) {
	err := clipboard.WriteAll(str)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
