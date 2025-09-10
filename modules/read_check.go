package modules

import (
	"errors"
	"fmt"
	"log"
	"os"
)

func Check() string {
	var name string

	fmt.Print("Print name of file u wanna check(to stop print \"stop\"): ")
	fmt.Scan(&name)

	_, err := os.Stat(name)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist")
	} else {
		fmt.Println("File exist")
	}

	return name
}

func Read() string {
	var name string

	fmt.Print("Print name of file u wanna to read(to stop print \"stop\"): ")
	fmt.Scan(&name)

	data, err := os.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))

	return name
}
