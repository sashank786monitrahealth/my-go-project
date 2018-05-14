package main

import (
  "fmt"
)

func myPrinter(w [9]string){
  for _,word := range w {
    fmt.Printf("%s ",word);
  }
  fmt.Printf("\n");

}

func main(){
  words := [9] string {"the","quick","brown","fox","jumped","over","the","lazy","log"};
  myPrinter(words);
}
