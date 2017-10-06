
/*Author : Vyacheslav Lukyanov
Id : G00339839
Program description :Newton’s method for square roots
*/



package main

import (
    "fmt"
    "math"
)

func newtSqrt(x float64) float64 {

    //declare variables
    var counter int
    z:=1.0
    var z_next float64

for  {

    counter++

     z_next = z - ((z*z - x) / (2 * z))//newton method sqr root


      if z_next == z{
        
           break
                 
            }

      fmt.Printf("Counter % d %f\n",counter,z_next)
      z=z_next
    }

    return z_next
}

func main() {

    var input int

    fmt.Println("Please enter a number to find a square root : ")

    fmt.Scan(&input)


    fmt.Println(newtSqrt(float64(input)))//newton method 

    fmt.Println(math.Sqrt(float64(input)))//math sqr root
}