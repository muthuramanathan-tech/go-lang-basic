package main

import (
	"time"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func say(s string){
	for i := 0; i<3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*100)
	}
	wg.Done()
}

func main(){
	//go say("Helllloo") // concurrency
	//go say("Helllloo")  
	//go say("Are u thr")
	
	// time.Sleep(time.Second) //this not best practice so willintroduce waitGroup

	wg.Add(1)
	go say("Hey")
	wg.Add(1)
	go say("Wovg")
	wg.Wait()
}