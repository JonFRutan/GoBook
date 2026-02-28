package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

//function that takes varying time
func printTask(id int, waitGroup *sync.WaitGroup) {
	//we use defer here to tell the function to wait to execute this until the very end of the task.
	//furthermore, defer will execute whether the process exits gracefully or errors out.
	defer waitGroup.Done() //tells the waitGroup that a process is done
	for i := 1; i <= 10; i++ {
		fmt.Printf("Goroutine %d: step %d\n", id, i)
		//wait between 0-2 seconds
		//time.Duration defaults to nano seconds, so we multiply it by time.Second
		waitTime := time.Duration(rand.Intn(3)) * time.Second
		time.Sleep(waitTime) 
	}
}

func main() {
	var waitGroup sync.WaitGroup
	fmt.Println("Starting tasks...")
	//we run routines independently
	//go is used to instantaite a new "goroutine", a very lightweight thread
	for i:=1;i<=10;i++{
		waitGroup.Add(1) //add to the waitGroup. This tells the waitgroup it has another process to wait to finish.
		go printTask(i, &waitGroup)
	}
	waitGroup.Wait()
	fmt.Println("Done!")
}