package main

import (
  "fmt"
)

func printer(msg1 string , msg2 string) error {
  _, myerr := fmt.Printf("msg1 = %s , msg2 = %s.  \n",msg1,msg2);
  return myerr;
}

func main() {

   errorMessage := printer("Hello World.","printer");
   fmt.Printf("errorMessage = %s\n\n", errorMessage);

}
