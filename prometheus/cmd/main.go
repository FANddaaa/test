package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/FANddaaa/test/prometheus"
)

func main() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)

	s := prometheus_test.NewServer("cmd/config.yaml")
	s.Init()
	s.Start()

	<-sc
	fmt.Println("receive signal to stop prometheus_test service")
	s.Stop()
}
