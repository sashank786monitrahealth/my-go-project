package main

import (
  "fmt"
  "math"
)

func halfIt(val *float64){
  *val = *val /2;
  fmt.Printf("halfIt() = %.5f\n", *val);
}

func main(){
  e := math.E;
  fmt.Printf("before halfIt() e = %.5f\n",e);
  halfIt(&e);
  fmt.Printf("before halfIt() e = %.5f\n",e);
}
