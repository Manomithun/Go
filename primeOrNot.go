package main

import "fmt"

func main(){
	var n int;
	fmt.Scan(&n)
	 b:=call(n)
	fmt.Println(b)
}
func call(n int) bool{
	if n<=1{
		return false
	}
	if n==2 {
		return true
	}
	 i:=2;
	for i*i<=n{
		if n%i==0{
			return false
		}
		i++
	}
	return true
}