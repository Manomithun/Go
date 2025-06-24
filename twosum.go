package main
import "fmt"

func main(){
	arr:=[5]int{2,7,9}
	target:=9
	mm:=make(map[int]int)
	for _,v:=range arr{
		com:=target-v
		if _,found:=mm[com]; found{
          fmt.Println(com,"+",v,"=",target)
		  break;
		}
		mm[v]=1

	}

}