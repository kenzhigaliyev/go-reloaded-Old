// package main

import (
	"fmt"
)

func ConvertBase(nbr, baseFrom, baseTo string) string {
	num1:=Atoi(nbr)
	str1:=PrintNbrBase(num1,baseFrom)
	num2:=Atoi(str1)
	return PrintNbrBase(num2,baseTo)
	
}

func PrintNbrBase(nbr int, base string) string{
	positive:=true
	result:=""
	if len(base)<2{
		return result
	}
	if nbr<0{
		positive=false
		nbr=nbr*-1
	}
	for index,char1:=range base{
		if char1=='-' || char1=='+'{
			return result
		}
		for _,char2:=range base[index+1:]{
			if char1==char2{
				return result
			}
		}
	}
	for nbr>len(base){
		val:=nbr%len(base)
		result=base[val:val+1]+result
		nbr=nbr/len(base)
	}
	result=base[nbr:nbr+1]+result
	if !positive{
		result="-"+result
	}
	return result
}

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
	result := ConvertBase("101011", "01", "0123456789")
	fmt.Println(result)
}
