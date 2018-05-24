package main;

import "fmt";

type myInput int;

func main(){
  var currVal int32;
  var prevVal int;
  var total int32 = int32(prevVal) + currVal;

  var inp myInput;
  var myTestInt int = int(inp);

  fmt.Println(total);
  fmt.Println(myTestInt);

var a rune;
var b string;
a = rune(69);
b = string(rune(a));
fmt.Printf("a = %v , b = %v\n",a , b);
fmt.Printf("datatype(a) = %T , datatype(b) = %T\n",a , b);

}
