package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	args := os.Args[1:]
	if len(args) <= 1 {
		fmt.Println("Usage: go run time.go a 1499875281")
		os.Exit(2)
	}
	switch args[0] {
	case "a":
		i, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			fmt.Println("Error")
		}
		fmt.Println(time.Unix(i, 0).Format("2006-01-02 15:04:05"))
	case "b":
	default:
		fmt.Println(args)
	}
