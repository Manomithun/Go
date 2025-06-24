package main 
import (
	"fmt"
	"sync"
)

const a int=100000
func main(){
     done:=make(chan struct{})
	  c:=0;
    mutex:=sync.Mutex{}
	 go func(){
         for i:=0;i<a;i++{
			mutex.Lock()
			c++;
			mutex.Unlock()
		 }
		 done<-struct{}{}
	 }()
	 go func(){
         for i:=0;i<a;i++{
			mutex.Lock()
			c--;
			mutex.Unlock()
		 }
		 done<-struct{}{}
	 }()
	 <-done
	 <-done
	 fmt.Println(c)
}