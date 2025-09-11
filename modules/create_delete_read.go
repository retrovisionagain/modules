package modules

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

func Create(name string) string {
	file, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	fmt.Println("File does not exist. Creating new file:", name)
	return name
}

func Read(name string) string {
	data, err := os.ReadFile(name)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File %s: \n", name)
	fmt.Println(string(data))

	return name
}

func Delete(name string) string {
	err := os.Remove(name)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist:", name)
	} else {
		fmt.Println("File deleted:", name)
	}

	return name
}

func Write() string {
	var name string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Print name of file u wanna to edit(to stop print \"stop\"): ")
	scanner.Scan()
	name = scanner.Text()

	if name == "stop" {
		fmt.Println("Goodbye!")
		return "stop"
	}

	_, err := os.Stat(name)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist:", name)
	}

	fmt.Print("Now edit this file: ")
	scanner.Scan()
	text := scanner.Text()

	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Printf("Failed to open file: %v", err)
		return ""
	}

	_, err = file.WriteString(text + "\n")

	if err != nil {
		log.Printf("Failed to write to file: %v", err)
		return ""
	}

	defer file.Close()
	fmt.Println("Done writing:", name)
	return name
}
