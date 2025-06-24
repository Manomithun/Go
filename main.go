package main

import (
	"fmt"
	//  "strings"
  // "math"
)
type studnt struct{
  name string
  age int
  marks map[string]int 
}
 func (s studnt) call() int{
    return s.age
   }
 func main(){
	//fmt.Println("hello, This is my first program in go")
     var nameone string="mano"
	 var nametwo ="Mano Mithun"
	 fmt.Println(nameone,nametwo)
	 //var nameone string="mano" ibnsted of decalaring like thse we can initisile
// 	greet:="Hello"
// 	fmt.Println(greet)
//     var age int=19
// 	var age1=19
// 	age2:=19
// 	fmt.Println("kumar's age ",age,"kav's age \n",age1,age2)
// 	fmt.Printf("hi this is %v and my age is %v \n",nameone,age)
// 	fmt.Printf("hi this is %q and my age is %q \n",nameone,age)
// 	fmt.Printf("hi this is %T and my age is %T \n",nameone,age)
// 	fmt.Printf("hi this is %v and my age is %f \n",nameone,18.6)
// 	fmt.Printf("hi this is %v and my age is %0.2f \n",nameone,2555.5)
//       var vv=fmt.Sprintf("hi this is %v and my age is %f \n",nameone,18.6);
// 	  fmt.Println(vv)
  
// // Arrays
//   var arr [3]int=[3]int{1,2,3}
//   fmt.Println(arr,len(arr))
//   var brr=[3]int{2,3,4}
//   fmt.Println(brr,len(brr))
//   crr:=[4]int{5,6,7}
//    fmt.Println(crr,len(crr))
//    names:=[5]string{"kavi","pavi","ravi","savi","Muni"}
//    fmt.Println(names,len(names))

// Slice
  // var score=[]int{1,2,3,4,5}
  // fmt.Println(score)
  // score[1]=10
  //  fmt.Println(score)
  //  range1:=score[0:3]
  //   fmt.Println(range1)
	//   range2:=score[:3]
  //   fmt.Println(range2)
	//   range3:=score[2:]
  //   fmt.Println(range3)

    // greet:="Hi Buddy"
    // fmt.Println(strings.Contains(greet,"Hi"));
    // fmt.Println(strings.Index(greet,"b"));
    // fmt.Println(strings.Index(greet,"B"));
    // fmt.Println(strings.ToLower(greet));
    // fmt.Println(strings.ToUpper(greet));
    // fmt.Println(strings.ReplaceAll(greet," ","*"));
    // fmt.Println(strings.Split(greet," "));
    // greeting:=strings.Split(greet," ");
    // fmt.Println(greeting);


  // Loop
  //  x:=0
  //  for x<5{
  //     fmt.Println(x)  
  //     x++
  //   }
  // for x:=0; x<=5 ;x++{
  //   fmt.Println(x);
  // }
  // ages:=[5]int{1,2,3,4,5}

  //   for index,value:= range ages{
  //     fmt.Println(index,value)
    
  // }


  // function

    //  greet("mano")
    //  greet("Mithun")
    //  peopleNames:=[]string{"mano","mithun","rasu","mani"};
    //  welcomeMess(peopleNames,greet);
    // circle:= getArea(4);
    //  fmt.Println(circle)
      // a:=10
      // if a>0 {fmt.Println(a)
      // }else if a<5{
      //   fmt.Println("a = %v",a)
      // }else{
      //   fmt.Println(" a is smaller")
      // }      

      // f1,f11:=first("Kavi Pavi")
      // fmt.Println(f1 ,f11);
      // b:="ManoMarutharasu"
      // fmt.Println(b[:10])

      // fmt.Println(TotalStudents);
      
  //     createmap:=map[string]int{
  //       "mano":19,
  //       "mithun":22,
  //       "marutharasu":50,
  //        "mani":40,
  //     }
  //     fmt.Println(createmap)
  //     fmt.Println(createmap["mano"])

  //     for K,V:=range createmap{
  //       fmt.Println(K,"-",V)
  //     }
  //     fmt.Println("HIIIIIIIIIIIII")
  //     createmap["ravi"]=37
  //      for K,V:=range createmap{
  //       fmt.Println(K,"-",V)
  //     }

  //  }   
  //  func first(n1 string ) (string,string){
  //  n:=strings.Split(n1," ")
  //  var ind[] string;
  //  for _,value:=range n{
  //      ind=append(ind,value[:1])
  //  }
  //  if(len(ind)>1){
  //   return ind[0],ind[1];
  //  }
  //  return ind[0],"_"
    

  kavi:=studnt{
     name:"kavi",
     age:17,
     marks: map[string]int{
      "tamil":90,
      "English":97,
     },
  }
  fmt.Println(kavi);
  fmt.Println(kavi.age)
  fmt.Println(kavi.call);
   }
  
    // func getArea(v float64) float64{
    //   return math.Pi * v * v
 //   }

    // func greet(name string){
    //   fmt.Println("Hello Bhai", name);
    // }
    //  func welcomeMess(name[] string,f func(name string)){
    //   for _,values:=range name{
    //     f(values);
    //   }
    //  }
