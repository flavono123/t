package main

import "os"

func main() {
	// Go >= 1.16
	// func ReadFile(filename string) ([]byte, error)
	bytes, err := os.ReadFile("input_os.txt")
	if err != nil {
		panic(err)
	}

	// func WriteFile(name string, data []byte, perm FileMode) error
	err = os.WriteFile("output_os.txt", bytes, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
