package main
import (
	"fmt"
	"time"
)
func main(){
	
	fmt.Println("Hi1")
	 go delayedPrint()
     go delayedPrint();
	fmt.Println("Hi")
	time.Sleep(3*time.Second)
}
func delayedPrint(){
	for i:=0;i<5;i++{
time.Sleep(1* time.Second)
	 fmt.Println("Hello",i)
	}
	 
}