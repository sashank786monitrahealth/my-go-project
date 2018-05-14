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
  words := [] string {"the","quick","brown","fox","jumped","over","the","lazy","log"};
  myPrinter(words[5:]);
  myPrinter(words[0:9]); // prints till the end of array
}
