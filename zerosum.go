package main
import "fmt"

func main(){
	arr:=[]int{4,2,-3,1,6}
	mm:=make(map[int]int)
	var pre int=0
	mm[0]=-1
	for i,v:=range arr{
		pre+=v;
	if _,found:=mm[pre]; found{
		fmt.Println(i-mm[v]+1)
		break
	}else{
     mm[pre]=i
	}
	
	}
	// delete(mm,1)
}