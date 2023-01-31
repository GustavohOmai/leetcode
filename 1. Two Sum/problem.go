package main

import "fmt"


func main(){
	arrayDeNumeros := [4]int{2, 7, 11, 15}
	target := 9
	arrayLength:=len(arrayDeNumeros)

	for i := 0; i < arrayLength; i++ {
		for j := 0; j < arrayLength; j++ {
			
			if  arrayDeNumeros[i]+arrayDeNumeros[j]==target {
				fmt.Println(i,j)
			}
		}
		
	}
}