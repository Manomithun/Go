package main
import (
	"fmt"
	"time"
)

func main(){
	done:=make(chan bool)
	go delayedPrint(done)
	<-done
}
func delayedPrint(done chan bool){
	time.Sleep(3*time.Second)
	fmt.Println("hi")
	done<-true
}