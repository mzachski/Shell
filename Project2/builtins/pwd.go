package builtins

import (
	"fmt"
	"os"
)

func PrintWorkingDirectory() {
	cwd, err := os.Getwd() //Get curr directory
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(cwd) //Print curr directory
}
