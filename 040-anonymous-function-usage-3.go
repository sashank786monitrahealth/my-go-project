package main

import (
  "fmt"
  "math"
)

func main(){
  for deg:=0.0; deg<=360 ; deg += 45.0 {


    var rad float64;

    rad = func() float64 {
      return deg * math.Pi /180;
    }();

    fmt.Printf("angle in degrees = %.2f ; angle in radians = %.2f\n",deg,rad);
  }
}
