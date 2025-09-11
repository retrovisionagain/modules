package main

import (
	"fmt"
	"smartgopher/modules"
)

// testing functions and logic of project
func main() {
	var date_now string
	// checking all modules
	fmt.Println("hey yo, it", modules.CheckTime_now(date_now))
	modules.Read(".gitignore")
	// delete.c moved to directory "test"
	modules.Delete("delete.c")
	// modules.Create("create.c") commented so as not to irritate
	modules.Write()
}
