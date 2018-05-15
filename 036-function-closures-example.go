package main

import (
  "fmt"
)

var (
   square = func(inp int) int {
     return multiply2nos(inp, inp);
   }

   multiply2nos = func(inp1, inp2 int) int {
     return inp1 * inp2
   }

)

func main(){
  fmt.Printf("multiply2nos(15,6) = %d\n",multiply2nos(15,6));
  fmt.Printf("square(36) = %d\n",square(36));
}
