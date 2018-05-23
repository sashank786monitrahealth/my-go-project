package main

import(
  "fmt";
)

func main(){
  var hadBreakFast bool = true;
  if !hadBreakFast {
    fmt.Printf("Let's Eat!\n");
  } else {
    fmt.Printf("It's Tea Time!\n")
  }
}
