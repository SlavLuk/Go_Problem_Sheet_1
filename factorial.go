/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Finding a factorial number !n and sum all digits in factorial
*/

package main

//import packages
import (
	"fmt"
	"math/big"
	"strconv"
)

func factorial(n int64) big.Int {

	if n == 0 {

		return *big.NewInt(1)
	}

	var f big.Int

	f.MulRange(1, n)

	return f

}

//start program with main function
func main() {

	//	variables declaration ant initialization
	num10 := big.NewInt(10)
	mod10 := big.NewInt(10)

	sum := big.NewInt(0)
	remain := big.NewInt(0)

	var total big.Int
	var input string

	fmt.Println("Please enter a number :") //promt for user input

	_, err := fmt.Scan(&input) //read user input

	if err != nil { //if error return

		fmt.Println("Error occured")

		return
	}

	fnum, err := strconv.ParseInt(input, 10, 64) //parse user input into int64

	if err != nil { //return if error

		fmt.Println(err) //print error

		return
	}

	fact := factorial(fnum) //invoke factorial method and assign to var fact

	sum, remain = sum.DivMod(&fact, num10, mod10) //divide by 10 and modulo  divide by 10 and assign to var

	total.Add(remain, big.NewInt(0)) //sum up remainder of modulo

	for sum.Cmp(big.NewInt(0)) != 0 { //loop and check on remaining if any digit left

		sum, remain = sum.DivMod(sum, num10, mod10) //same as above

		total.Add(remain, &total) //same as above

	}

	fmt.Printf("The Factorial number of !%d is : %d\n", fnum, &fact) //print out factorial number
	fmt.Printf("The Total sum of digits : %d\n", &total)             //print out sum of digits in factorial

}
