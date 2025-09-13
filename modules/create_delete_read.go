package modules

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Create(name string) string {
	file, err := os.Create(name)

	if err != nil {
		log.Println("Cant create file:", err)
	}

	defer file.Close()
	fmt.Println("File does not exist. Creating new file:", name)
	return name
}

func Read(name string) string {
	data, err := os.ReadFile(name)

	if err != nil {
		log.Println("Cant read file:", err)
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

func Write(filename, text string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Println("Error while modifying file:", err)
	}
	defer file.Close()

	_, err = file.WriteString(text + "\n")
	if err != nil {
		log.Println("Error while modifying file:", err)
	}

	fmt.Println("Done writing:", filename)
	return nil
}

func Mkdir(dir string) {
	parent_dir := filepath.Dir(dir)
	info, err := os.Stat(parent_dir)

	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("Root directory not exist:", parent_dir)
		}
		return
	}

	if !info.IsDir() {
		fmt.Println("Path is correct,but it not directory:", parent_dir)
		return
	}

	err = os.MkdirAll(dir, 0750)

	if err != nil {
		log.Println("Failed to create new directory:", err)
		return
	}

	fmt.Println("Directory has succesfull created")
}
