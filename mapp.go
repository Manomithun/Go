package main
import "fmt"
func main(){
	var n int;
	fmt.Scan(&n)
	arr := make([]int, n)
	i:=0
	for i<n{
		fmt.Scan(&arr[i])
		i++
	}
	mmap := make(map[int]int)
	i=0
	for i<n{
		mmap[arr[i]]=mmap[arr[i]]+1;
		i++
	}
	for K,V:=range mmap{
		fmt.Println(K,":",V)
	}
     max:=0
	 ans:=0
	 for K,V:=range mmap{
		if max<V {
			max=V
			ans=K
		}
	}
    fmt.Println(ans)

}