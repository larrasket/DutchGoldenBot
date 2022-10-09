package main

import (
	"fmt"
	"github.com/kkdai/twitter"
	"net/http"
	"os"
	"sync"

	"github.com/dghubble/sling"
)

func rain() {
	client := &http.Client{}
	req, _ := sling.New().Get("https://example.com/").Request()
	x, f := client.Do(req)
	fmt.Println(x.StatusCode)
	fmt.Println(f)
}
func xmain() {
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Print("1")
	}()

	go func() {
		defer wg.Done()
		fmt.Print("2")
	}()

	go func() {
		defer wg.Done()
		fmt.Print("3")
	}()
	wg.Wait()
}

func mainr() {
	hello := func(wg *sync.WaitGroup, id int) {
		defer wg.Done()
		fmt.Printf("Hello from %v!\n", id)
	}
	const numGreeters = 5
	var wg sync.WaitGroup
	wg.Add(numGreeters)
	for i := 0; i < numGreeters; i++ {
		go hello(&wg, i+1)
	}
	wg.Wait()
}

func mainxs() {
	var count int
	var lock sync.Mutex
	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}
	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}
	// Increment
	var arithmetic sync.WaitGroup
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			increment()
		}()
	}

	// Decrement
	for i := 0; i <= 5; i++ {
		arithmetic.Add(1)
		go func() {
			defer arithmetic.Done()
			decrement()
		}()
	}
	arithmetic.Wait()
	fmt.Println("Arithmetic complete.")

}

func lmain() {
	defer fmt.Println("eff")
	defer fmt.Println("fff")
	defer fmt.Println("ffzzzzz")
}

func xmaine() {

	os.Setenv("GOTWI_API_KEY", "q8KPEGtgCI8yFoiyAonmMvm9s")
	os.Setenv("GOTWI_API_KEY_SECRET", "CNgHkbYAceoUErwFr1RiSyOGO6msdWQ77TDiZGdiDKZCjz3J61")

}

const (
	//Get consumer key and secret from https://dev.twitter.com/apps/new
	ConsumerKey    string = "q8KPEGtgCI8yFoiyAonmMvm9s"
	ConsumerSecret string = "CNgHkbYAceoUErwFr1RiSyOGO6msdWQ77TDiZGdiDKZCjz3J61"
)

func main() {
	twitterClient = NewDesktopClient(ConsumerKey, ConsumerSecret)

	//Show a UI to display URL.
	//Please go to this URL to get code to continue
	twitterClient.DoAuth()

	//Get timeline only latest one
	timeline, byteData, err := twitterClient.QueryTimeLine(1)

	if err == nil {
		fmt.Println("timeline struct=", timeline, " byteData=", string(byteData))
	}
}
