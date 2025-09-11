package modules

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
)

// TODO: переписать все функции из read_check и из этого файла,а также из time так,чтобы функции принимали на вход только название файла и по возможности доп. параметры
func Create() string {
	var name string

	fmt.Print("Print name of file u wanna to create or write(to stop print \"stop\"): ")
	fmt.Scan(&name)

	file, err := os.Create(name)

	if err != nil {
		log.Fatal(err)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist. Creating new file.")
	}
	fmt.Print("File was created: ")
	defer file.Close()
	return name
}

func Delete() {
	var name string

	fmt.Print("Print name of file u wanna to delete(to stop print \"stop\"): ")
	fmt.Scan(&name)

	if name == "stop" {
		fmt.Println("Goodbye!")
	}

	err := os.Remove(name)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist:", name)
	}
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("File deleted:", name)
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
		fmt.Println("File does not exist. Creating new file.")
	}

	fmt.Print("Now edit this file: ")
	scanner.Scan()
	text := scanner.Text()

	err = os.WriteFile(name, []byte(text), 0644)
	if err != nil {
		log.Printf("Failed to write file %s: %v", name, err)
		return ""
	}

	fmt.Print("Done writing: ")
	return name
}
