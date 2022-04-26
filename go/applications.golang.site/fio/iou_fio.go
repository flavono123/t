package main

import (
	"io/ioutil"
	"os"
)

func main() {
	// func ReadFile(filename string) ([]byte, error)
	// - As of Go 1.16, this function simply calls os.ReadFile.
	bytes, err := ioutil.ReadFile("input_iou.txt")
	if err != nil {
		panic(err)
	}

	// func WriteFile(name string, data []byte, perm FileMode) error
	// https://errorsingo.com/compliation-no-new-variables-on-left-side-of/
	// type FileMode = fs.FileMode
	err = ioutil.WriteFile("output_iou.txt", bytes, os.FileMode(0644))
	if err != nil {
		panic(err)
	}
}
