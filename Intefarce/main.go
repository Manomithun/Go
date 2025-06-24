package main

import (
	"fmt"
	"reflect"
	"time"
	"os"
	
)
type logger interface{
	Info(string)
}
type screenlog struct{}
func (screenlog) Info(mess string){
	fmt.Println("%v:info:%s",time.Now(),mess)
}
type FileLogger struct{
	File *os.File
}
func (l *FileLogger) Info(mess string){
	fmt.Fprintf(l.File,"%v:info:%s",time.Now(),mess)
}
func NewFileLogger(filename string)* FileLogger{
	file,err:=os.OpenFile(filename,os.O_CREATE|os.O_APPEND|os.O_WRONLY,0777);
	if(err!=nil){
		panic(err)
	}
	return &FileLogger{file}
}
func main(){
    var log logger
	// log=&screenlog{}
	log=NewFileLogger("Output.txt")
	log.Info("Hello")
	fmt.Println(reflect.TypeOf(log))
}