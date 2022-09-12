package tasks
import "fmt"


func Task1(){
	fmt.Println("Fizz Buzz Program")
	for count:= 1; count<=100; count++ {
		fmt.Print(count)
		fmt.Print(": ")
		if(count%3==0){
			fmt.Print("Fizz ")
		}
		if(count%5==0){
			fmt.Print("Buzz")
		}
		fmt.Println("")
	}
}