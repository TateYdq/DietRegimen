package ai

import (
	"fmt"
	"time"
)

func Task() {
	t := time.NewTicker(3*time.Second)
	defer t.Stop()
	fmt.Println(time.Now())
	time.Sleep(4*time.Second)
	for {
		select {
		case <-t.C: fmt.Println("6666666")
		}
	}
}
