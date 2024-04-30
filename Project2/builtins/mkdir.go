package builtins

import (
	"os"
)

func MakeDirectory(directory string) error {
	err := os.Mkdir(directory, 0755)
	if err != nil {
		return err
	}
	return nil
}
