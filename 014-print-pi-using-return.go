package main

import (
    "fmt"
    "math"
)

func getPi() float64{
    return math.Pi;
}

func main(){
  fmt.Printf("The value of Pi = %v\n",getPi());
}
