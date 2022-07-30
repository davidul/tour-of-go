package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}

	printDir(dir, 0)

	file, err := os.OpenFile("tmp", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(file)

	b := []byte("Hello there")
	n, err := file.Write(b)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Bytes written %d \n", n)

	openFile, err := os.OpenFile("tmp", os.O_RDONLY, 0644)
	defer func(openFile *os.File) {
		err := openFile.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(openFile)

	readBytes := make([]byte, 1024)
	read, err := openFile.Read(readBytes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(read)
	//bytes := readBytes[:]
	fmt.Printf("Read bytes %s", string(readBytes))

	/*for _, file := range dir{
		fmt.Println(file.Name())
		if file.IsDir() {

		}
	}*/

}

func printDir(dir []os.DirEntry, lvl int) {
	lvl += 1
	for _, file := range dir {
		tabs := ""
		for i := 0; i < lvl; i++ {
			tabs += "\t"
		}
		fmt.Println(tabs, file.Name())
		if file.IsDir() {
			dirs, _ := os.ReadDir(file.Name())
			printDir(dirs, lvl)
		}
	}
}
