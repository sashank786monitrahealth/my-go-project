package main;

import "fmt";

func main(){
  intPtr := new(int);
  *intPtr =44;

  p := new(struct{first, last string});
  p.first = "Basavaraj";
  p.last = "Badiger";

  fmt.Printf("Value = %d, type = %T\n", *intPtr, intPtr);
  fmt.Printf("Person = %+v type = %T\n",*p, p);
}
