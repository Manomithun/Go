package main

import "fmt"
func main(){
	var n int
	fmt.Scan(&n)
	i:=2
	for n%i==0{
		fmt.Println(i)
		n=n/i
	}
	for i=3; i*i<=n;i+=2{
		if n%i==0{
fmt.Println(i)
		n=n/i
		}
		
	}
	if n>2{
	  fmt.Println(n);
	}
}