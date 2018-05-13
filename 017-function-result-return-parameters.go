package main

import (
    "fmt"
    //"bytes"
)

func add(inp0 int, inp1 int) (int) {
    return inp0+inp1;
}

func sub(inp0 int, inp1 int) (int){
    return inp0 - inp1;
}
//type myString string;


func main(){

var myAdd func(int, int) int = add;
mySub := sub;

fmt.Printf("myAdd(36,45) = %d\n",myAdd(36,45));
fmt.Printf("mySub(36,45) = %d\n",mySub(36,45));

}
