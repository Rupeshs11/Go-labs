package main

import (
	"fmt"
	"sync"
)

// take a example (suppose we are creating the social media post)
type post struct {
	views int
	mu    sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	//defer wg.Done()

	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()
	
	p.mu.Lock()
	p.views += 1
}

// mutex is used for prevent from race condition
func main() {
	var wg sync.WaitGroup
	myPost := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go myPost.inc(&wg)
	}
	wg.Wait()
	fmt.Println(myPost.views)
}
