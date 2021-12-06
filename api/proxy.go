package api

import (
	"log"
	"sync"
)

func Proxy(username string, password string, address string, wg *sync.WaitGroup) {
	cookies := login(username, password)
	name := check(cookies, username)
	fmtstr := ""
	if len(name) == len("夏驰") {
		fmtstr = "%v-%v   %v\n"
	} else {
		fmtstr = "%v-%v %v\n"
	}
	log.Printf(fmtstr, username, name, submit(cookies, address))
	wg.Done()
}
