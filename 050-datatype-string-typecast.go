package main;

import(
   "fmt";
)

func main(){
   var myStr1 string;
   var myStr2 string;
   myStr1 = fmt.Sprintf("%v",16.04);
   myStr2 = "My Ubuntu Version : "+myStr1+"\n";
   fmt.Printf(myStr2);
   fmt.Printf("length of above string = %d\n",len(myStr2));
}
