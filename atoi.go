package main

import (
	"fmt"
)

func Atoi(s string) int{
	length:=len(s)
	if length == 0{
		return 0
	}
	result:=0
	positive:=true
	number:=[]rune{}
	sign:= 1
	for _,char:=range s{
		number=append(number,char)
	}
	if number[0]=='-'{
		for i:=1;i<length;i++{
			if number[i]<'0' || number[i]>'9'{
				return 0
			}
		}
		sign=-1
		positive=false
	} else if number[0]=='+' || (number[0]>='0' && number[0]<='9'){
		for i:=1;i<length;i++{
			if number[i]<'0' || number[i]>'9'{
				return 0
			}
		}
	} else {
		return 0
	}
	for i:=0;i<length;i++{
		if number[i]>='1' && number[i]<='9'{
			if (length-i)>19{
				return 0
			}
			for j:=i;j<length;j++{
				result=(result*10) + (int(number[j])-48)
			}
			result=sign *result
			if result<0 && positive{
				return 0
			} else if result>0 && !positive{
				return 0 
			} else {
				return result
			}
		}
	}
	return result
}

func main() {
	fmt.Println(Atoi("12345"))
	fmt.Println(Atoi("0000000012345"))
	fmt.Println(Atoi("012 345"))
	fmt.Println(Atoi("Hello World!"))
	fmt.Println(Atoi("+1234"))
	fmt.Println(Atoi("-1234"))
	fmt.Println(Atoi("++1234"))
	fmt.Println(Atoi("--1234"))
	fmt.Println(Atoi("-2147483648"))
	fmt.Println(Atoi("19223372036854775808"))
	fmt.Println(Atoi("-009223372036854775809"))
}
