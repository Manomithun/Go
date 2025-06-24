package main 
import "fmt"
func main(){
	ch:=make(chan int)
	go call(ch)
	
	for a:=range ch{
		fmt.Println(a);;
	}
	
}
func call(out chan int){
	for i:=0;i<5;i++{
		out<-i
	}
	close(out);
}