package main

import(
   "fmt";
)

var (
   char1 = 'ϕ';
   char2 = 'అ';
   char3 = 'ఆ';
   char4 = 'ఇ';
   char5 = 'ఈ';
   char6 = 'ఉ';
   char7 = '\b';
   char8 = '\t';
   char9 = '\n';
   char10 = '\u0369';
   char11 = '\xFA';
   char12 = '\045';
)

func main(){
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char1,char1,char1);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char2,char2,char2);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char3,char3,char3);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char4,char4,char4);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char5,char5,char5);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char6,char6,char6);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char7,char7,char7);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char8,char8,char8);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char9,char9,char9);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char10,char10,char10);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char11,char11,char11);
   fmt.Printf("The character = %c The int32-number = %v and datatype = %T\n",char12,char12,char12);
}
