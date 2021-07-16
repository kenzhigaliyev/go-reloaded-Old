package main

import (
	"fmt"
)

func Split(s, sep string) []string {
	var str = []string{} 
	new:=""
	check:=""
	length:=len(sep)
	separ:=false
	count:=0
	for index,char:=range s{
		if char == 'H'{
			for index1,val := range s[index:(length+index)]{
				fmt.Println(string(val))
				if (index1==(length-1)) {
					check=check+string(val)
					if check==sep{
						// fmt.Println(check)
						check=""
						separ=true
						if len(new)>0{
							str=append(str,new)
						}
						new=""
						// fmt.Println(1)
						break
					} else {
						// fmt.Println(check)
						check=""
						separ=false
						new=new+string(char)
						// fmt.Println(2)
						break
					}
					// fmt.Println(check)
					break
				} else {
					check=check+string(val)
					// fmt.Println(check)
				}
			}
			fmt.Println(separ)
		} else if separ==true{
			count++
			if count==(length-1){
				separ=false
				count=0
			}
		} else if index ==(len(s)-1){
			new=new+string(char)
			fmt.Println(new)
			str=append(str,new)
		} else {
			new=new+string(char)
			fmt.Println(new)

		}
	}
	return str
}

func main() {
	s := "HAHelloHAhowHAareHAyou?"
	fmt.Println(Split(s, "HA"))
}
