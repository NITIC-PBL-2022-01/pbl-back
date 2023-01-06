package main

import (
	"os"
	"os/signal"
	"syscall"
)

func main() {
  sig := make(chan os.Signal, 1)
  signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

  <-sig
}
