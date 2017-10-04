/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Finding out if a word is palindrome
*/

package main

import	(
	"fmt"
	"strings"
)

func palindromeTest(s string)bool{
	
	var reversed[]rune//declare slice

	s = strings.ToLower(s)//convert to lower case
	
	origin:=[]rune(s)//convert a string into slice char
	
	reversed =make([]rune,len(origin))//make slice
		
	for i,j :=len(origin)-1,0;i>=0;i,j = i-1,j+1 {//reverse string
	
	reversed[j]=origin[i]
		
	}
	
	return string(reversed)==string(origin)//compare two strings
}

func main(){

	var input string//declare string var

	fmt.Println("Please enter a string : ")

	fmt.Scan(&input)//read user input

	if	palindromeTest(input) {//check the input

		fmt.Println("Entered word is palindrome")

	}else{

		fmt.Println("Entered word is not palindrome")

	}

	

}