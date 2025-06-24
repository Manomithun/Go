package main

import ("fmt"
// "bufio"
// "os"
)

func main(){
   
//    reader:=bufio.NewReader(os.Stdin);
//   n,_:=reader.ReadString('\n')
   var n int;
   fmt.Scan(&n)
	fmt.Println(n)
	i:=1
	for i<n{
		if n%i==0{
			fmt.Println(i)
		}
		i++
	}

}