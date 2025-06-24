package main

import "fmt"

func main(){
	var n int;
	fmt.Scan(&n)
	arr:=make([]int,n)
	i:=0
	for i<n{
		fmt.Scan(&arr[i])
		i++
	}
	fmt.Println(arr)
	var nn int=int(n/2)
	i=0;
	j:=nn-1;
	for i<j{
		var t int=arr[i]
		arr[i]=arr[j]
		arr[j]=t
		i++
		j--
	}
	fmt.Println(arr)
}