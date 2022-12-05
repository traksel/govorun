package main

import (
	project_pkg "govorun/pkg"
	"time"
)

func main() {
	project_pkg.InitLogFile()
	for i := 0; i <= 999; i++ {
		project_pkg.RaiseErrorLog("Error #", i)
		time.Sleep(3 * time.Second)
		project_pkg.RaiseWarnLog("Warn #", i)
		time.Sleep(3 * time.Second)
		project_pkg.RaiseInfoLog("INFO #", i)
		time.Sleep(3 * time.Second)
		if i == 999 {
			i = -1
		}
	}
}
