package main

import (
	"fmt"
	// "time"
)

func someFunc(num string){
	fmt.Println(num)
}

func main() {
	//This is go routines
	// go someFunc("3")
	// go someFunc("2")
	// go someFunc("1")

	// time.Sleep(time.Second * 2)

	// fmt.Println("Hello")

	//chaneels

	// myChannel := make(chan string)
	// anotherChannel := make(chan string)

	// go func(){
	// 	//syntax for sync data to channel
	// 	myChannel <- "data"
	// }()

	// go func(){
	// 	//syntax for sync data to channel
	// 	anotherChannel <- "cow"
	// }()
	
	// select{
	// case msgFromMyChannel := <-myChannel:
	// 	fmt.Println(msgFromMyChannel)
	
	// case msgFromanotherChannel := <-anotherChannel:
	// 	fmt.Println(msgFromanotherChannel)

	// }
	

	//For-Select-Loop

	charChannel := make(chan string,3)
	chars := []string{"a","b","c"}

	for _,s := range chars{
		select{
		case charChannel <-s:
		}
	}
	close(charChannel)

	for result := range charChannel{
		fmt.Println(result)
	}
}