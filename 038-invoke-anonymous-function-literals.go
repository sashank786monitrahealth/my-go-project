package main

import "fmt"

func main(){
  fmt.Printf(
    "%d (F) = %.2f (C) \n", 0,
    func (f float64)  float64 {
        myInp := f;
        myCel := (myInp - 32.0) * (5.0/9.0);
        return  myCel;
    }(0.0),
    );
}
