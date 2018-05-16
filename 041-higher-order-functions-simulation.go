package main

import "fmt"

func apply(numbersInp []int, inputFunction func(int) int )func(){
  for i, v := range numbersInp {
      numbersInp[i] = inputFunction(v);
  }
  return func() {
    fmt.Println(numbersInp);
  }
}

func main(){
     myNumbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
     result := apply(myNumbers, func(i int) int {
        return i*2;
     })
     result();
}
