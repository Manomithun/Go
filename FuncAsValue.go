package main 
import "fmt"
func main(){
	fn:=func (a int)int{
		return a*a
	}
	var n int
	fmt.Scanf("%d",&n)
	fmt.Println(fn(n))
	fmt.Println(sq(n,fn))
	k:=newFunc(n)
	var input int
	fmt.Print("Enter number")
	fmt.Scan(&input);
	ans:=k(input)
	fmt.Println(ans)
}
func sq(a int,fn func(a int)int)int{
	return fn(a);
}
func newFunc(n int)func(int)int{
	if n==2{
	return func(a int) int{
		return a
	}
}else{
	return func(a int) int{
		return a*a
	}
}
}