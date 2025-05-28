package main

import (
	"log"
	"os"
)

func main(){
	if len(os.Args) < 2{
		log.Fatal("Expected subcommand: [list, describe]")
	}

	args := os.Args[2:]
	switch os.Args[1] {
	case "list":
		listArgs, err := GetListArgs(args)
		if err != nil {
			log.Fatal("Error parsing list arguments: ", err)
		}
		log.Println(listArgs)
	case "describe":
		log.Println(args)
		describeArgs, err := GetDescribeArgs(args)
		if err != nil {
			log.Fatal("Error parsing describe args: ", err)
		}
		log.Printf("DescribeArgs: %+v\n", describeArgs)
	}
}
