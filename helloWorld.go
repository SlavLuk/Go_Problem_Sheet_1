/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Printing Hello World in Japanese characters
*/

package main

import "fmt" //import package

//start program with main function
func main() {

	const HELLO_WORLD = "こんにちは世界" //declare const

	fmt.Println(HELLO_WORLD) //print Hello World

	for _, code := range HELLO_WORLD {

		fmt.Printf("%#U\n", code) //print unicode for Japanese characters

	}

}
