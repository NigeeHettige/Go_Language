package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main(){
	fmt.Println("Hey Welcome to my quiz game")
	// var name float64 = 2.0
	// var name bool =false
	// var name string = "Tim"

	// var name string = "Tim"
	// name = "Joe"
	// fmt.Println(name)

	// name := "Tim"
	// age := 21
	// fmt.Printf("Hello %v , You are %v?",name,age)
	// var name string
	// var age int
	// fmt.Scan(&name)
	// fmt.Println(name)
	// fmt.Scan(&age)
	// fmt.Println(age)
	fmt.Printf("Enter your name: ")
	var name string
	fmt.Scan(&name)

	fmt.Printf("Hello, %v, welcome to the game! \n",name)

	fmt.Printf("Enter your age: ")
	var age int
	fmt.Scan(&age)

	if(age>=10){
		fmt.Println("You are eligible")
	}else{
		fmt.Println("You are not eligible")
		return
	
	}
	fmt.Println("continue")
	
	
	
	getline := bufio.NewReader(os.Stdin)
	getline.ReadString('\n')
	fmt.Printf("What is better ,The RTX 3080 or RTX 3090?")
	answer, _ := getline.ReadString('\n')
	answer = strings.TrimSpace(answer)

	if(answer == "RTX 3080"){
		fmt.Println("You are correct")
	}else{
		fmt.Println("Icorrect answer")
	}
}