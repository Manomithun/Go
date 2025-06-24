package main 
import "fmt"

type cels float32
type fah float32

func (k cels) toF()fah{
	return fah(9.0/5.0*k+32)
}

func main(){
	var k cels
	k=3.32
	fmt.Println(k.toF())
}
