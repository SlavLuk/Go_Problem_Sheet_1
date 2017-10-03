
/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Finding the largest and the smallest numbers of a list
*/

package main

import	"fmt"

func bigAndSmall(l[]int)(int,int){//function finding largest and smallest numbers

	lar :=0//declare and initialize to 0
	
	for _,num:=range l{//go thru every elements of the list
		
		if lar<num{
		
			lar=num
		
			}
					
		}

	small := lar//declare and initialize to largest number in the list

	for _,sm:=range l{//go thru all elements in the list

		if small >sm{
				
			small=sm
				
			}

		}

		return lar,small//return two values
}

func main(){


list:=[]int{4,2,3,5,16,100000,55,23,100}//slice list


largest,smallest := bigAndSmall(list)//invoke and assign returned values

//print outputs
fmt.Printf("The largest number is %d\n",largest)
fmt.Printf("The smallest number is %d\n",smallest)

}