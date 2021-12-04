package api

import (
	"log"
	"sync"
)

func Proxy(username string, password string, wg *sync.WaitGroup) {
	cookies := login(username, password)
	name, err := check(cookies)
	if err != nil {
		panic(err)
	}
	if len(name) == len("夏驰") {
		log.Printf("%v-%v   %v\n", username, name, submit(cookies))
	} else {
		log.Printf("%v-%v %v\n", username, name, submit(cookies))
	}
	wg.Done()
}
