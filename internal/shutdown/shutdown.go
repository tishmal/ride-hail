package shutdown

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func GracefulStop(stopFunc func(context.Context) error, timeout time.Duration) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	<-stop
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	_ = stopFunc(ctx)
}
