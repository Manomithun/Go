package main
import "fmt"

func main(){
	var n int;
	fmt.Scan(&n)
	arr:=make([]int,n)
	brr:=make([]int,n)
	for i:=0;i<n;i++{
		fmt.Scan(&arr[i])
	}
	for i:=0;i<n;i++{
		fmt.Scan(&brr[i])
	}

	// mm:=make(map[int]int)
	// for i:=0;i<n;i++{
	// 	mm[arr[i]]++
	// }
	// var ans int;
	// for _,V:=range arr{
    //     if(mm[V]==1){
	// 		ans=V;
	// 		break;
	// 	}
	// }
	// fmt.Println(ans)
	//itersection
	mm:=make(map[int]bool)
    for _,v:=range arr{
		mm[v]=true
	}
    seen:=make(map[int]bool)
	res:=[]int{}
	for _,v:=range brr{
		if mm[v] && !seen[v]{
			res=append(res, v)
			seen[v]=true
		}
	}
	fmt.Println(res)
	for _,V:=range res{
		fmt.Println(V)
	}

}