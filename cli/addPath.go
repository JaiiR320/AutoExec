package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readPath() string {
	fmt.Print("Enter file path: ")
	var filePath string
	reader := bufio.NewReader(os.Stdin)
	filePath, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err, "\n", "Read: ", filePath)
		return ""
	}
	filePath = strings.Replace(filePath, "\"", "", -1)
	return strings.TrimSpace(filePath)
}

func validPath(filePath string) bool {
	if filePath == "" {
		return false
	}
	_, err := os.Stat(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
		} else {
			fmt.Println("Error checking file:", err)
		}
		return false
	}
	return true
}

func AddPath(bucket string) {
	var filePath string
	for {
		filePath = readPath()
		if validPath(filePath) == true {
			break
		}
	}

	file, err := CreateBucket(bucket)
	if err != nil {
		fmt.Println("Error fetching bucket:", err)
		return
	}
	defer file.Close()

	_, err = fmt.Fprintln(file, filePath)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	fmt.Println("Added program to bucket.")
}
