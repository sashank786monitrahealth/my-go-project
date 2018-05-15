package main

import (
  "fmt"
  "math"
)

func doubleIt(val float64){
  val = 2 * val;
  fmt.Printf("doubleIt() = %.5f\n", val);
}

func main(){
  e := math.E;
  fmt.Printf("before doubleIt() e = %.5f\n",e);
  doubleIt(e);
  fmt.Printf("before doubleIt() e = %.5f\n",e);
}
