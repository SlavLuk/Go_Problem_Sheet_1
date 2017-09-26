/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Current date and time
*/

package main

//import packages
import (
	"fmt"
	"time"
)

//start program with main function
func main() {

	timeNow := time.Now() //returns current local time

	fmt.Printf("The Date %10d/%d/%d\nTime now %10d:%d:%d", timeNow.Day(), timeNow.Month(),
		timeNow.Year(), timeNow.Hour(), timeNow.Minute(), timeNow.Second())

}
