/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Prints numbers from 1 to 100 and Fizz ,Buzz in case of multiples by 3 or 5
*/

package main

//import packages
import (
	"fmt"
)

//start program with main function
func main() {

	//loop from 1 to 100
	for i := 1; i <= 100; i++ {

		if i%3 == 0 && i%5 == 0 { //check if a number may be divided by 3 and 5

			fmt.Println("FizzBuzz")

		} else if i%3 == 0 { //check if a number may be divided by 3

			fmt.Println("Fizz")

		} else if i%5 == 0 { //check if a number may be divided by 5

			fmt.Println("Buzz")

		} else {

			fmt.Println(i) //print out numbers

		}
	}
}
