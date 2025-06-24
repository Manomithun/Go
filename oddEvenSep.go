package main 
import "fmt"
func  main(){
	input:=make(chan int)
	evenSquare:=make(chan int)
	merge:=make(chan int)
	print:=make(chan int)
	go getInput(input)
	go mergeit(input,evenSquare,merge)
	go printit(print,evenSquare,merge)
	for a:=range print{
		fmt.Println(a)
	}
}
func getInput(input chan int){
	for i:=0;i<10;i++{
		input<-i;
	}
	close(input);
}
func mergeit(input chan int, evenSquare chan int, merge chan int){
	for a:=range input{
		if a%2==0{
			evenSquare<-a;
		}else{
            merge<-a;
		}
	} 
	close(evenSquare)
	close(merge)
}
func printit(print chan int, evenSquare chan int, merge chan int) {
	done := make(chan bool)

	go func() {
		for a := range evenSquare {
			print <- a
		}
		done <- true
	}()

	go func() {
		for a := range merge {
			print <- a
		}
		done <- true
	}()
	<-done
	<-done
	close(print)
}