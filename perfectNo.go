package main
import "fmt"

func main(){
	var n int
	fmt.Scan(&n)
	var sum int=0
	for i:=1;i<=n/2;i++{
		if n%i==0{
          sum+=i
		} 
	}
	if sum==n{
		fmt.Println("Yes")
	}else{
		fmt.Print("No")
	}
}