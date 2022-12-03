package main

import (
	"time"

	"github.com/traksel/logovoril/modules"
)

func main() {
	for i := 999; i <= 999; i++ {
		modules.RaiseErrorLog("Error #", i)
		time.Sleep(10 * time.Second)
		if i == 999 {
			i = -1
		}
	}
}
