package main

import (
	"go-base/files"
)

func main() {
	//imt.IMT()
	//finance.Finance()
	//bookmarks.Bookmarks()
	// password.Password()

	files.ReadFile()
	files.WriteFile("Hello", "hi.txt")
}
