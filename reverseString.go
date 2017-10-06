/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Reversing string
*/

package main

import	(
	"fmt"
)

func reversedString(s string)string{
	
	var reversed[]rune//declare slice
	
	origin:=[]rune(s)//convert a string into slice char
	
	reversed =make([]rune,len(origin))//make slice
		
	for i,j :=len(origin)-1,0;i>=0;i,j = i-1,j+1 {
	
	reversed[j]=origin[i]
		
	}
	
	return string(reversed)//cast and return reverse string
}

func main(){

	var input string//declare string var

	fmt.Println("Please enter a string : ")

	fmt.Scan(&input)//read user input

    revStr:=reversedString(input)
	

	fmt.Printf("The original string is %s\nThe reverse string is %s\n",input,revStr)

}