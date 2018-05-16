package main

import (
  "fmt"
  "math"
)

func main(){
  for i:=0.0; i<=360 ; i += 45.0 {
    fmt.Printf("angle in degrees = %.2f ; angle in radians = %.2f\n",i,
    func() float64 {
      var rad float64;
      rad = i * math.Pi /180;
      return rad;
    }(),
    )
  }
}
