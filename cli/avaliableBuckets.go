package cli

import (
	"fmt"
	"os"
	"strings"
)

func GetBuckets() []string {
	files, err := os.ReadDir(bucketPath)
	if err != nil {
		err = os.Mkdir(bucketPath, 0777)
		if err != nil {
			fmt.Println("Error creating bucket directory:", err)
			return nil
		}
		files, err = os.ReadDir(bucketPath)
	}

	if len(files) == 0 {
		fmt.Println("No buckets found.")
		return nil
	}

	var buckets []string
	for _, file := range files {
		name := file.Name()
		name = strings.TrimSuffix(name, ".txt")
		buckets = append(buckets, name)
	}
	return buckets
}

func AvaliableBuckets() {
	buckets := GetBuckets()
	if buckets == nil {
		return
	}
	for _, bucket := range buckets {
		fmt.Println("  " + bucket)
	}
}

func RunAutoExec() {
	buckets := GetBuckets()
	if buckets == nil {
		return
	}
	for _, bucket := range buckets {
		if bucket == "autoexec" {
			RunBucket(bucket)
			return
		}
	}
}
