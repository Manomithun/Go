package main 
import "fmt"
func main(){
	calll(10,30)
	calll(40.5,44.7)
	fmt.Println(maaxx(2,7),maaxx(8.9,8.7))
}
func calll[T any](x,y T){
	fmt.Println(x,y);
}
func maaxx[T int | int32 |int64 | float64](x,y T)T{
	if x<y {
		return y
	}
		return x
	
}