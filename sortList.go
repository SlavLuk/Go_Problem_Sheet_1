/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Merging and sorting two slices into one
This programme only merging and sorting integer type
*/

package main

import(

	"fmt"
	"sort"
)

func mergeAndSort(listA[]int,listB[] int)([]int){

for _,i:=range listB{

	listA=append(listA,i)//append element from slice listB to the end of slice A

}

return listA//return new slice two combined

}//end merge func

func main(){

	//declare two slices
a:=[]int{1,2,5,6,13}
b:=[]int{3,7,9}

c:=mergeAndSort(a,b)//invoke merge func and assign returned slice to a new slice
sort.Ints(c)//sort slice of integer

fmt.Printf("%v",c)//print out values in slice 


}