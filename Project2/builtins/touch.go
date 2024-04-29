package builtins

import (
	"log"
	"os"
)

func TouchFile(filename string) error {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	log.Printf("File Touched.")

	return nil
}
