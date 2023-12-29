package cli

import (
	"fmt"
	"os"
)

func ReadBucket(bucket string) {
	content, err := os.ReadFile(bucketPath + bucket + ".txt")
	if err != nil {
		fmt.Println("That bucket does not exist")
		return
	}
	fmt.Print(string(content))
}
