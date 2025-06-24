package main
import "fmt"
func main(){
	i:=1
	n:=5
	// for i<n{
	// 	for k:=1;k<(n-i);k++{
	// 		fmt.Print(" ")
	// 	}
	// 	for j:=0;j<i+(i-1);j++{
	// 		fmt.Print("*")
	// 	}
	// 	fmt.Println()
	// 	i++
	// }
	for i<n{
		for k:=1;k<(n-i);k++{
		fmt.Print(" ")
        }
		for j:=0;j<i;j++{
			fmt.Print("* ")
		}
		fmt.Println()
		i++
	}
}