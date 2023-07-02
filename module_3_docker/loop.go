package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	flag.Parse()
	glog.V(2).Info("生成两个死循环线程")
	glog.Flush()
	go func() {
		for {

		}
	}()

	for {
		v := 1
		glog.V(2).Info(v)
	}
}
