package runtime

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func SetupSignal(cancel context.CancelFunc) {
	go func() {
		termChan := make(chan os.Signal, 1)
		signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
		<-termChan
		cancel()
	}()
}
