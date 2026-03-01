package main

import (
	"fmt"
	"time"
	"sync"
	"math/rand"
)

type result struct {
	Id int
	Sum int
	IsOdd bool
}

//function that takes varying time
//The first set of parentheses indicates the necessary parameters to pass to this function
//The second set of parentheses indicates a return value, this allows us to return with just the statement `return`
//this is know as a "naked" return
func printTask(id int, waitGroup *sync.WaitGroup, taskChannel chan *result) (taskSum int, isOdd bool){
	//we use defer here to tell the function to wait to execute this until the very end of the task.
	//furthermore, defer will execute whether the process exits gracefully or errors out.
	defer waitGroup.Done() //tells the waitGroup that a process is done
	for i := 1; i <= 10; i++ {
		var rnumber = rand.Intn(10)
		fmt.Printf("Goroutine %d: step %d - Adding %d to sum.\n", id, i, rnumber)
		taskSum += rnumber    //add the random number to the local task sum
		//fmt.Printf("Number Grabbed! ID: %d\n", id)
		//wait between 0-2 seconds
		//time.Duration defaults to nano seconds, so we multiply it by time.Second
		waitTime := time.Duration(rand.Intn(2)) * time.Second
		time.Sleep(waitTime) 
	}
	if taskSum % 2 == 0 {
		isOdd = false
	} else {
		isOdd = true
	}
	taskChannel <- &result{Id: id, Sum: taskSum, IsOdd: isOdd}
	//%d is used for a decimal number, %s for strings, %t for bool values.
	fmt.Printf("!-- TASK %d FINISHED WITH SUM OF %d IS IT ODD? %t --!\n", id, taskSum, isOdd)
	return
}

func main() {
	//creating a waitgroup that will make the main process wait for all goroutines to finish before continuing
	var waitGroup sync.WaitGroup
	//channels are used for inter routine data communication
	//the channel below is unbuffered. Unbuffered channels will cause channel inputs (from our printTask function) to actually freeze until the number is accepted.
	//unbuffered channels enforce synchrony with this, they are *similar* to a buffered channel with a buffer size of 1 - however the goroutines won't freeze unless the channel is unbuffered.
	//we can create a buffered channel like so:
  //sumChannel := make(chan int 10)
	taskChannel := make(chan *result)
	var tasks []*result 
	var printSum int
	fmt.Println("Starting tasks...")
	//we run routines independently
	//go is used to instantaite a new "goroutine", a very lightweight thread
	for i:=1;i<=50;i++{
		waitGroup.Add(1) //add to the waitGroup. This tells the waitgroup it has another process to wait to finish.
		go printTask(i, &waitGroup, taskChannel)
	}
	//anonymous goroutine to wait and close the sumchannel.
	//this will prevent the main process from blocking on waitGroup.wait()
	go func() {
		waitGroup.Wait()
		close(taskChannel)
	}()
	
	//read from taskChannel until the channel is closed
	for task := range taskChannel {
		printSum += task.Sum
		tasks = append(tasks, task)
	}
	//after the channel is closed by our anonymous function, we'll summarize our results
	var oddCount int
	var evenCount int 
	for _, result := range tasks {
		if result.IsOdd {
			oddCount += 1
		} else {
			evenCount += 1
		}
	}
	average := printSum / len(tasks)
	fmt.Printf("Done with with sum - %d. We had %d odds, %d evens, and an average of %d per task.\n", printSum, oddCount, evenCount, average)
}