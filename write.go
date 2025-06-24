package main
import (
	"os"
	"io"
	"fmt"
)
func main(){
	// data:=[]byte("hell Guys")
	// err:=os.WriteFile("output.txt",data,0644)
	// if err!=nil{
	// 	panic(err)
	// }
     var bytes[10]byte
	f,err:=os.Open("output.txt")
	if err!=nil{
		panic(err)
	}
	for {
		_,err:=f.Read(bytes[:])
		if err==io.EOF{
			break;
		}
		
	}
	fmt.Printf("%s",bytes)
	f.Close()
}