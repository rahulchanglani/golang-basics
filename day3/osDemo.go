package main

import (
	"fmt"
	"os"
)

func main() {
	// create file

	// fp, _ := os.Create("newfile.txt")
	fp, _ := os.Create("D:\\SDE\\go\\go-lang-training\\day3\\newfile.txt")
	fmt.Println("file created", fp)

	// create folder
	fd := os.Mkdir("myfolder",0644) // windows permission to create folder 0777, rwx

	// 0644 is octal mode, RW is for owning user
	fmt.Println("folder created", fd)
}