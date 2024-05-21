// Objective:
// Learn how to send and receive values using channels.

// 1)Task:

// Write a program that creates a goroutine to send a message "Hello, World!" to a channel.
// The main function should receive the message from the channel and print it.
package main

import (
	"fmt"
	"strings"
)

func sendMessage(ch chan string) {
	ch <- "Hello, World!"
}

func main() {
	ch := make(chan string)
	go sendMessage(ch)
	fmt.Println(<-ch)

}

// Hints:

// Use an unbuffered channel for simplicity.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to create and use goroutines.

// 2)Task:

// Write a program that launches three goroutines. Each goroutine should print numbers from 1 to 5 with a delay of 1 second between each number.
// Ensure that the main function waits for all goroutines to finish before exiting.
// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main
import (
    "fmt"
    "sync"
    "time"
)
func PrintNumbers(wg *sync.WaitGroup){
    for i:=1;i<=5;i++{
        fmt.Println(i)
        time.Sleep(1* time.Second)
    }
    wg.Done()
}


func main() {
    var wg sync.WaitGroup
    wg.Add(3)
    
    for i:=0;i<3;i++{
        go PrintNumbers(&wg)
    }
    wg.Wait()
}
//to know how the working of add(),wait(),Done()
package main
import (
	"fmt"
	"time"
)
type waitgroup struct{
    count int
}
func (wg *waitgroup) Add(num int){
    wg.count+=num
}
func (wg *waitgroup) Done(){
    wg.count--
}
func (wg *waitgroup) Wait(){
    for wg.count!=0{
        time.Sleep(time.Second)
    }
}
func loop(message string,wg *waitgroup){
    fmt.Println(message)
    wg.Done()
}

func main(){
	var wg waitgroup
	wg.Add(3)

	go loop("message1",&wg)
	go loop("message2",&wg)
	go loop("message3",&wg)
	
	wg.Wait()

}
// Hints:

// Use time.Sleep for delays.
// Use a sync.WaitGroup to wait for all goroutines to complete.

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to handle multiple senders with a single receiver using channels.

// 3)Task:

// Write a program where two goroutines send different messages ("Hello" and "World") to the same channel.
// The main function should receive both messages from the channel and print them.
package main

import (
	"fmt"
)

func sendMessage(msg string, ch chan string){
    ch <-msg
}

func main() {
    ch:=make(chan string)
    go sendMessage("Hello",ch)
    go sendMessage("World",ch)
    
    msg1:= <-ch
    msg2:= <-ch
    
    fmt.Println(msg1,msg2)
}

// -------------------------------------------------------------------------------------

// Objective:
// Understand how to use channels for communication between goroutines.

// 4)Task:

// Write a program where the main function creates a goroutine that generates numbers from 1 to 10 and sends them to a channel.
// The main function should receive these numbers from the channel and print them.
// Hints:
package main

import (
	"fmt"
)

func PrintMessage(num chan int){
    for i:=1;i<=10;i++{
        num <- i
    }
    close(num)
}

func main() {
    num:=make(chan int)
    
    go PrintMessage(num)
    
   for n:=range num{
       fmt.Println(n)
   }
   }


// Use an unbuffered channel for simplicity.
// Remember to close the channel when done sending values.

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use a buffered channel.

// 5)Task:

// Write a program that creates a buffered channel with a capacity of 3.
// The main function should send three integers (1, 2, 3) to the buffered channel without blocking.
// Then, receive and print the integers from the channel.
package main

import (
	"fmt"
)

func Capacity(numbers []int,ch chan int){
    for _,num:=range numbers{
        ch <- num
    }
}

func main() {
    m:=make(chan int,3)
    
    numbers := []int{1,2,3}
    go Capacity(numbers,m)
    
   for i:=0;i<len(numbers);i++ {
       num := <-m
       fmt.Println(num)
   }
}

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// 6)Task:

// Write a program that launches two goroutines. Each goroutine sends a series of messages to a channel.
// The main function should use a select statement to receive messages from both channels and print them.
// Hints:
package main
import (
	"fmt"
	"time"
)
func SeriesMessage(message []string,ch chan string){
	for _,msg:=range message{
		ch <-msg
		time.Sleep(time.Second)
	}
}
func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)

	message1:=[]string{"helo","rinshad","hii"}
	message2:=[]string{"hi","rrrr","good"}

	go SeriesMessage(message1,ch1)
	go SeriesMessage(message2,ch2)


	for i:=0;i<2;i++{
		select{
		case msg:= <-ch1:
			fmt.Println("channel 1 is:",msg)
		case msg:= <-ch2:
			fmt.Println("channel 2 is:",msg)	
		default:
			fmt.Println("no message is received")	
			time.Sleep(time.Second)
		}
	}
}
// Use two separate channels.
// Use the select statement inside a loop to receive from the channels.

// -------------------------------------------------------------------------------------

// Objective:
// Use sync.WaitGroup to wait for multiple goroutines to complete.

// 7)Task:

// Write a program that launches three goroutines, each printing a different message.
// Use a sync.WaitGroup to ensure the main function waits for all goroutines to finish before exiting.
package main
import (
	"fmt"
// 	"time"
	"sync"
)

func DiffMessage(message string,wg *sync.WaitGroup){
    fmt.Println(message)
    // time.Sleep(time.Second)
    wg.Done()
}

func main(){
	var wg sync.WaitGroup
	wg.Add(3)
	
    go DiffMessage("message1",&wg)
     go DiffMessage("message2",&wg)
      go DiffMessage("message3",&wg)
    
	
	wg.Wait()

}

// -------------------------------------------------------------------------------------

// Objective:
// Learn how to use the select statement to handle multiple channels.

// 8)Task:

// Write a program that creates two channels and two goroutines. Each goroutine sends a different message to its respective channel.
// Use a select statement in the main function to receive messages from both channels and print them.
package main
import (
    "fmt"
    )
    func MultiChannel(message string,ch chan string){
        ch <-message
    }
    func main(){
        ch1:=make(chan string)
        ch2:=make(chan string)
        
        message1:= "Hello from rinshad"
         message2:= "Hello from rashhhh"
         
         go MultiChannel(message1,ch1)
         go MultiChannel(message2,ch2)
       for i:=0;i<2;i++{ 
        select{
            case msg:=<-ch1:
            fmt.Println("the message 1 is :",msg)
             case msg:=<-ch2:
            fmt.Println("the message 2 is :",msg)
          
        }
       }

    }