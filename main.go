package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Wrong number of arguments")
		return
	}

	if len(os.Args) == 1 {
		fmt.Println("C:/Users/sakari.ekqvist/source/side/terminalshortcuts/conf.txt")
		return
	}

	key := os.Args[1]
	if key == "edit" {
		cmd := exec.Command("cmd", "/C start C:/Users/sakari.ekqvist/source/side/terminalshortcuts/conf.txt")
		err := cmd.Start()
		if err != nil {
			log.Fatal(err)
		}
		err = cmd.Wait()
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	file, err := os.Open("C:/Users/sakari.ekqvist/source/side/terminalshortcuts/conf.txt")
	if err != nil {
		fmt.Println("Failed to open file:", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ";")
		if splits[0] == key {
			fmt.Println(splits[1])
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Failed to read file:", err)
		return
	}
}
