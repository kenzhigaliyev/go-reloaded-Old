package main

import (
	"os"
	// "fmt"
)

func Operation(arg string) bool {
	if len(arg) == 1 && ((arg == "+") || (arg == "-") || (arg == "*") || (arg == "/") || (arg == "%")){
		return true
	}
	return false
}

func Calculate(num1, num2 int, arg string) (int,int,bool){
	result:=0
	val:=0
	positive:=true
	sign:=false
	if arg == "+" {
		if num1<0 && num2<0{
			positive=false
		}
		result=num1+num2
		if (positive && result<0) || (!positive && result>0){
			return 0,val,sign
		}
	} else if arg == "-"{
		if num1<num2{
			positive=false
		}
		result=num1-num2
		if ((num1!=result+num2) || (positive && result<0)) || ((num1!=result+num2) || (!positive && result>0)){
			return 0,val,sign
		}
	} else if arg == "*"{
		result=num1*num2
		if num1!=result/num2{
			return 0,val,sign
		}
	} else if arg == "/"{
		if num2==0{
			val=val+1
			return result,val,sign
		}
		result=num1/num2
	} else {
		if num2==0{
			val=val+2
			return result,val,sign
		}
		result=num1%num2
	}
	if result<0{
		result=-1*result
		sign=true
	}
	return result,val,sign
}

func Atoi(s string) (int,bool) {
	length:=len(s)
	if length == 0{
		return 0, false
	}
	result:=0
	positive:=true
	number:=[]rune{}
	sign:=1
	for _,char:=range s{
		number=append(number,char)
	}
	if number[0]=='-'{
		for i:=1;i<length;i++{
			if number[i]<'0' || number[i]>'9'{
				return 0,false
			}
		}
		sign=-1
		positive=false
	} else if number[0]=='+' || (number[0]>='0' && number[0]<='9'){
		for i:=1;i<length;i++{
			if number[i]<'0' || number[i]>'9'{
				return 0, false
			}
		}
	} else {
		return 0, false
	}
	for i:=0;i<length;i++{
		if number[i]>='1' && number[i]<='9'{
			if (length-i)>19{
				return 0, false
			}
			for j:=i;j<length;j++{
				result=(result*10) + (int(number[j])-48)
			}
			result=sign *result
			if result<0 && positive{
				return 0, false
			} else if result>0 && !positive{
				return 0, false
			} else {
				return result, true
			}
		}
	}
	return result, true
}

func IntToStr(number int,sign bool) string{
	new:=""
	for number>10{
		result:=number%10
		number=number/10
		new=string(result+48)+new
	}
	new=string(number+48)+new
	if sign {
		new="-"+new
	}
	return new
}

func main() {
	arg := os.Args[1:]
	if len(arg) != 3 {
		return
	}
	number1,val1:=Atoi(arg[0])
	number2,val2:=Atoi(arg[2])
	if (val1 && val2) && Operation(arg[1]){
		result,warn,sign:=Calculate(number1,number2,arg[1])
		if warn==1{
			os.Stdout.WriteString("No division by 0\n")
			return
		} else if warn==2{
			os.Stdout.WriteString("No modulo by 0\n")
			return
		}
		os.Stdout.WriteString(IntToStr(result,sign))	
		os.Stdout.WriteString("\n")
		return
	}
	os.Stdout.WriteString("0\n")
}
