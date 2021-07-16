package main

import (
	"fmt"
)

func AtoiBase(s string, base string) int {
	str:=[]rune{}
	result:=0
	pow:=0
	for index,char1:=range base{
		if char1 == '-' || char1 =='+'{
			return 0
		}
		for _,char2:=range base[index+1:]{
			if char1==char2{
				return 0
			}
		}
		str=append(str,char1)
	}
	for _,val:=range s{
		for i:=0;i<len(base);i++{
			if val==str[i]{
				fmt.Println(result)
				pow=i
				for i!=0{
					pow=pow*len(base)
					i--
				}
				result=result+pow
			}
		}
	}
	return result
}

// func AtoiBase(s string, base string) int {
// 	lengthBase := len(base)
// 	result := 0
// 	for _,char:= range base{
// 		if char=='-' || char=='+'{
// 			return result
// 		}
// 	}
// 	if lengthBase == 2 {
// 		for _, char := range base {
// 			if (char == '0') || (char == '1') {
// 			} else {
// 				return 0
// 			}
// 		}
// 		result = FromBinary(base, s)
// 	} else {
// 		lengthString:=len(s)
// 		for _, char1 := range s {
// 			for index2, char2 := range base {
// 				if char1 == char2 {
// 					lengthString--
// 					result = result + (index2*Power(lengthBase, lengthString))
// 				}
// 			}
// 		}
// 	}
// 	return result
// }

// func FromBinary(str string, number string) int {
// 	binary := 0
// 	pow := len(number) - 1
// 	for _, char := range number {
// 		if char == '1' {
// 			binary = binary + Power(2, pow)
// 			pow--
// 		} else if char == '0' {
// 			binary = binary + Power(0, pow)
// 			pow--
// 		}
// 	}
// 	return binary
// }

// func Power(number int, pow int) int {
// 	if pow != 0 {
// 		pow = pow - 1
// 		number = number * Power(number, pow)
// 	} else if pow == 0 {
// 		return 1
// 	}
// 	// fmt.Println(number)
// 	return number
// }


func main() {
	fmt.Println(AtoiBase("125", "0123456789"))
	fmt.Println(AtoiBase("1111101", "01"))
	fmt.Println(AtoiBase("7D", "0123456789ABCDEF"))
	fmt.Println(AtoiBase("uoi", "choumi"))
	fmt.Println(AtoiBase("bbbbbab", "-ab"))
}
