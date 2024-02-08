package main

import "fmt"

func doSomething(){
	fmt.Println("Doing something");
}

func printNumber (n int){
	fmt.Println(n);
}


func main(){
	doSomething(); // this will not fork a new thread, it will block the main thread until doSomething() is done
	go doSomething(); // this will fork a new thread and the main thread will continue to execute equiivalent in node.js is fork a new process by doing `child_process.fork()` or `child_process.spawn()`


	// currently the main thread is not in sync with the go routine
	// the main thread will exit before the go routine is done executing so we need to wait for the go routine to finish
	// we can use a channel to wait for the go routine to finish
	done := make(chan bool);
	go func(){
		doSomething();
		done <- true; // this will signal the main thread that the go routine is done executing
	}();
	<-done; // this will block the main thread until the go routine is done executing if we don't do this the main thread will exit before the go routine is done executing


	fmt.Println("Hello World");
}