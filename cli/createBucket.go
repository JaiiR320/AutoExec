package cli

import (
	"os"
)

func CreateBucket(bucket string) (*os.File, error) {
	file, err := os.OpenFile(bucketPath+bucket+".txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	return file, err
}
