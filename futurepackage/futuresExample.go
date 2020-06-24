package main

import (
	"fmt"
	"time"
	"futurepackage/futures"
)


func doSomething() futures.Result {
	time.Sleep(1 * time.Second)
	return futures.Result{Value: "exampleValue"}
}

func printOutput(value futures.Result){
	if value.Error != nil{
		fmt.Println("got error", value.Error)
	} else {
		fmt.Println("Result : ", value.Value)
	}
}

func main(){
	x:= futures.MakeFuture(doSomething) 
	go func(){
		printOutput(x.Get())
	}()
	go func(){
		printOutput(x.Get())
	}()
	go func(){
		printOutput(x.GetWithTimeout(500 * time.Millisecond))
	}()
	time.Sleep(1100 * time.Millisecond)
	x.Cancel()
	time.Sleep(3 * time.Second)
}
