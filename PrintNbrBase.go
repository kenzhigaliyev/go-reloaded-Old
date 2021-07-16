// package main

import (
	"github.com/01-edu/z01"
)


func PrintNbrBase(nbr int, base string){
	positive:=true
	result:=""
	if len(base)<2{
		return
	}
	if nbr<0{
		positive=false
		nbr=nbr*-1
	}
	for index,char1:=range base{
		if char1=='-' || char1=='+'{
			z01.PrintRune('N')
			z01.PrintRune('V')
			return
		}
		for _,char2:=range base[index+1:]{
			if char1==char2{
				z01.PrintRune('N')
				z01.PrintRune('V')
				return
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
		z01.PrintRune('-')
	}
	for _,char3:=range result{
		z01.PrintRune(char3)
	}
	return
}


func main() {
	PrintNbrBase(125, "0123456789")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "01")
	z01.PrintRune('\n')
	PrintNbrBase(125, "0123456789ABCDEF")
	z01.PrintRune('\n')
	PrintNbrBase(-125, "choumi")
	z01.PrintRune('\n')
	PrintNbrBase(125, "aa")
	z01.PrintRune('\n')
}
