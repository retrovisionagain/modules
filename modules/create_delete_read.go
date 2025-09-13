package modules

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Create(filename string) string {
	file, err := os.Create(filename)

	if err != nil {
		log.Println("Cant create file:", err)
	}

	defer file.Close()
	fmt.Println("File does not exist. Creating new file:", filename)
	return filename
}

func Read(filename string) string {
	data, err := os.ReadFile(filename)

	if err != nil {
		log.Println("Cant read file:", err)
	}

	fmt.Printf("File %s: \n", filename)
	fmt.Println(string(data))
	return filename
}

func Delete(filename string) string {
	err := os.Remove(filename)

	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("File does not exist:", filename)
	} else {
		fmt.Println("File deleted:", filename)
	}

	return filename
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
