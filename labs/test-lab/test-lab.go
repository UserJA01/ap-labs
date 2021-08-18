package main

import (
	"fmt"
	"os"
)

func main() {
	//fmt.Println("Welcome to the jungle")
	//fmt.Println("Welcome to the jungle", os.Args[1])
	totalInput := len(os.Args)
	if(totalInput < 2){
		fmt.Println("Please write down at least one name")
	}else{
		fmt.Print("Welcome to the jungle ")
		for i := 1; i < totalInput; i++ {
			fmt.Print(os.Args[i])
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
