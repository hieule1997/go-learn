package main
 
import (
	"time"
	"fmt"
	"sync"
)
var wg sync.WaitGroup
func say (s string){
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(s)
		time.Sleep(time.Millisecond*50)
	}
}
func main(){
	wg.Add(1)
	go say("helo")
	wg.Add(1)
	go say("hieu")
	wg.Wait()

	// time.Sleep(time.Second)
}