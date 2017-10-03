package main

import (//import packages
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func getUserInput()(int,error){//read user input and parse string into int

	var input string

	fmt.Println("Please enter your guess number between 1 to 100 :")//promt for input
	
		 fmt.Scan(&input)
	
		 num, err := strconv.Atoi(input) //parse user input into int
			  
		 return num,err
		
}

func main() {//start programme here

	//variables declaration
	var counter =1
	var prevNum int

	r := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(100)//generate random number between 1 and 100

	gnum,err:=getUserInput()//invoke function to read input

	if err!=nil{//check on errors

		fmt.Println(err)
		
		return
	}
	

	prevNum = gnum//assign to first user input num

	for r !=gnum {//loop until random num is == user guess num

		counter++//count inputs

		if gnum<r {

			fmt.Println("Your number is too low")

		}else{

			fmt.Println("Your number is too high")
		}

		gnum,err =getUserInput()//read user input
		
			if err != nil { //return if error
		
				fmt.Println(err) //print error
		
				return
			}

			if prevNum ==gnum {//check on repeated numbers

				counter--

			}
		
				prevNum = gnum

	}//end for loop

	//print out outputs
	fmt.Println("Well done you have guessed!")
	fmt.Printf("Your guessed number is %d\nYou have done %d tries\n",gnum,counter)
	fmt.Printf("Random number is %d\n",r)

}
