package main

import (
	"encoding/json"
	"fmt"
	"time"
)
type customer struct{
	id int `json:id`
	name string `json:name`
	dob time.Time `json:dob`
}
type transcation struct{
	id int `json:id`
	dob time.Time `json:dob`
	Amount int `json:Amount`
}
type Account struct{
	AccNo int `json:AccNO`
	Holder *customer "json:Holder"
	Transcation []transcation `json:Transcation`
}
func main(){
	acc:=Account{
		AccNo: 1,
		Holder: &customer{
			id:1,
			name: "mano",
			dob:time.Now(),
		},
		Transcation: []transcation{
			{id:1,dob:time.Now(),Amount: 100},
			{id:2,dob:time.Now(),Amount: 200},
			{id:3,dob:time.Now(),Amount: 300},

		},
	}
	data,err:=json.Marshal(acc)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("%s",data)
}