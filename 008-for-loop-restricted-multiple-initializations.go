package main

import (
  "fmt"
)

func main() {


   for myCounter,counter2 :=0, 1; myCounter < 10; myCounter, counter2 = myCounter+1 , counter2*2 {
     fmt.Printf("%d Hey There! %d \n",myCounter,counter2);

   }

}
