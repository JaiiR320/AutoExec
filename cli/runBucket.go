package cli

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func RunBucket(bucket string) {
	content, err := os.ReadFile(bucketPath + bucket + ".txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	for _, line := range strings.Split(string(content), "\n") {
		if line == "\n" || line == "" {
			continue
		}
		cmd := exec.Command("cmd", "/c", line)
		err := cmd.Start()
		if err != nil {
			fmt.Println("Error running command:", err)
			return
		}
		fmt.Println("Running ", line)
	}
}
