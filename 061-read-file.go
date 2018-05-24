package main;

import (
   "fmt";
   "io/ioutil"
)

func main(){
   b, err := ioutil.ReadFile("my-sentence.txt");
   if err != nil{
      fmt.Printf("Error-Encountered:: %v Error-DataType = %T\n",err,err);
   }
   fmt.Printf("\nBytes Output %v -- Bytes datatype = %T\n",b,b); // printing the contents as bytes
   fmt.Printf("\n"+string(b)+"\n"); // print the content as a string
}
