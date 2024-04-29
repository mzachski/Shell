package builtins

import (
	"fmt"
	"io"
	"os"
)

func Concatenate(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	// Print the contents to standard output
	fmt.Println(string(contents))

	return nil
}
