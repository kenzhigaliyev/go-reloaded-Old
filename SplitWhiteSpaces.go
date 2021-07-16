package main

import (
	"fmt"
)

func SplitWhiteSpaces(s string) []string {
	var str = []string{} 
	new:=""
	for index,char:=range s{
		if char == ' '{
			if new!=""{
				fmt.Println(new)
				str=append(str,new)
				new=""
			}
		} else if index ==(len(s)-1){
			new=new+string(char)
			str=append(str,new)
		} else {
			new=new+string(char)
		}
	}
	return str
}

func main() {
	fmt.Println(SplitWhiteSpaces(" d    Hello  how are   you?  "))
}
