package main

import (
  "fmt"
)

func myPrinter(w []string){
  for _,word := range w {
    fmt.Printf("%s ",word);
  }
  fmt.Printf("\n");

  w[3] = "elephant"; // this does not change the value in the array

}

func main(){
  words := make([]string,9)
  fmt.Printf("len = %d , cap = %d\n\n", len(words), cap(words));
  words[0] = "the";
  words[1] = "quick";
  words[2] = "brown";
  words[3] = "fox";
  words[4] = "jumps";
  words[5] = "over";
  words[6] = "the";
  words[7] = "lazy";
  words[8] = "dog";
  fmt.Printf("len = %d , cap = %d\n\n", len(words), cap(words));
  myPrinter(words[0:4]);
  myPrinter(words[3:9]); // prints till the end of array
}
