package cli

import "fmt"

func PrintHelp() {
	fmt.Println("USAGE:")
	fmt.Println("    list            - List avaliable buckets")
	fmt.Println("    create <bucket> - Create a new bucket")
	fmt.Println("    add    <bucket> - Add a path to a bucket")
	fmt.Println("    read   <bucket> - Read the contents of a bucket")
	fmt.Println("    run    <bucket> - Run the contents of a bucket")
	fmt.Println("    help            - Display this message")
}
