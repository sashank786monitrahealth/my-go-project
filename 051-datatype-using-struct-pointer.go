package main;

import(
  "fmt";
)



func main(){

myStructPtr := &struct{x ,y int} {22,44};
arrayPtr := &[2] string{"Agra","Bhubhaneswar"};

fmt.Printf("struct=%#v, type=%T\n",myStructPtr, myStructPtr);
fmt.Printf("arrayPtr=%#v, type=%T\n",arrayPtr,arrayPtr);
}
