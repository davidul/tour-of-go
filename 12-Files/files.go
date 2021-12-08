package main

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.ReadDir(".")
	if err != nil{
		panic(err)
	}

	printDir(dir, 0)

	/*for _, file := range dir{
		fmt.Println(file.Name())
		if file.IsDir() {

		}
	}*/

}

func printDir(dir[] os.DirEntry, lvl int){
	lvl += 1
	for _, file := range dir{
		tabs := ""
		for i := 0; i < lvl; i++{
			tabs += "\t"
		}
		fmt.Println(tabs,file.Name())
		if file.IsDir() {
			dirs, _ := os.ReadDir(file.Name())
			printDir(dirs, lvl)
		}
	}
}