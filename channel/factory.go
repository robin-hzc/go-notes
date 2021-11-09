package channel

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type Factory struct {
	wg     sync.WaitGroup
	worker chan bool
	msgCh  chan struct{}
	stop bool
}

func NewFactory() *Factory {
	return &Factory{
		wg:     sync.WaitGroup{},
		worker: make(chan bool, 2),
		msgCh:  make(chan struct{}, 2),
	}
}

func (a *Factory) Producer() {
	var n = 0
	go func() {
		for true {
			if a.stop {
				break
			}
			a.worker <- true
			go func() {
				a.wg.Add(1)
				//work time
				time.Sleep(10 * time.Second)
				<-a.worker
				a.msgCh <- struct{}{}
				n++
				log.Println("work..." + strconv.Itoa(n))
			}()
		}
	}()
}

func (a *Factory) Consumer() {
	var res struct{}
	go func() {
		for true {
			select {
			case res = <-a.msgCh:
				a.wg.Done()
			case <-time.After(15 * time.Second):
				log.Println("超时了...")
			}
		}
	}()
}

func (a *Factory) Stop() {
	defer func() {
		close(a.worker)
		log.Println("服务关闭...")
	}()
	a.stop = true
	a.wg.Wait()
}
