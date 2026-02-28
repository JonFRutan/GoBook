package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

//function that takes varying time
func printTask(id int, waitGroup *sync.WaitGroup, sumChannel chan int) {
	//we use defer here to tell the function to wait to execute this until the very end of the task.
	//furthermore, defer will execute whether the process exits gracefully or errors out.
	defer waitGroup.Done() //tells the waitGroup that a process is done
	for i := 1; i <= 100; i++ {
		var rnumber = rand.Intn(10)
		fmt.Printf("Goroutine %d: step %d - Adding %d to sum.\n", id, i, rnumber)
		sumChannel <- rnumber //we add the random number into our passed channel
		fmt.Printf("Number Grabbed! ID: %d\n", id)
		//wait between 0-2 seconds
		//time.Duration defaults to nano seconds, so we multiply it by time.Second
		waitTime := time.Duration(rand.Intn(10)) * time.Second
		time.Sleep(waitTime) 
	}
}

func main() {
	//creating a waitgroup that will make the main process wait for all goroutines to finish before continuing
	var waitGroup sync.WaitGroup
	//channels are used for inter routine data communication
	//the channel below is unbuffered. Unbuffered channels will cause channel inputs (from our printTask function) to actually freeze until the number is accepted.
	//unbuffered channels enforce synchrony with this, they are *similar* to a buffered channel with a buffer size of 1 - however the goroutines won't freeze unless the channel is unbuffered.
	//we can create a buffered channel like so:
  //sumChannel := make(chan int 10)
	sumChannel := make(chan int)
	var printSum int
	fmt.Println("Starting tasks...")
	//we run routines independently
	//go is used to instantaite a new "goroutine", a very lightweight thread
	for i:=1;i<=50;i++{
		waitGroup.Add(1) //add to the waitGroup. This tells the waitgroup it has another process to wait to finish.
		go printTask(i, &waitGroup, sumChannel)
	}
	//anonymous goroutine to wait and close the sumchannel.
	//this will prevent the main process from blocking on waitGroup.wait()
	go func() {
		waitGroup.Wait()
		close(sumChannel)
	}()
	
	for val := range sumChannel {
		printSum += val
	}
	fmt.Printf("Done with waitGroup1 with sum - %d\n", printSum)
}