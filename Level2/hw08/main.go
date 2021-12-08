// This is simple utility for finding duplicates.
package main

import (
	"flag"
	"fmt"
	"green/search"
	"os"
)

var (
	help bool
	dir  string
	del  bool

	dirFlag  = flag.String("dir", ".", "use dot for all scanning, or chose directory")
	delFlag  = flag.Bool("del", false, "use it fot delete duplicate files.")
	heplFlag = flag.Bool("help", false, "some help info")
)

func main() {
	flag.Parse()

	help = *heplFlag
	dir = *dirFlag
	del = *delFlag

	if help == true {
		fmt.Println("This is a utility for finding duplicates.\nYou can call it without flags, or specify the path to the directory via: -dir <directory path>")
		fmt.Println("If you want delete duplicates, use flag: -del\n ")
		os.Exit(1)
	}

	results := search.Scan(dir, del)
	for duplicate := range results {
		fmt.Println("Found duplicate: ", duplicate)
	}
}
