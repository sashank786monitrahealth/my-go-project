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
  words := make([]string,0,9)
  fmt.Printf("len = %d , cap = %d\n\n", len(words), cap(words));
  words = append(words,"the");
  words = append(words,"quick");
  words = append(words,"brown");
  words = append(words,"fox");
  words = append(words,"jumps");
  words = append(words,"over");
  words = append(words,"the");
  words = append(words,"lazy");
  words = append(words,"dog");
  fmt.Printf("len = %d , cap = %d\n\n", len(words), cap(words));
  myPrinter(words[0:4]);
  myPrinter(words[3:9]); // prints till the end of array
}
