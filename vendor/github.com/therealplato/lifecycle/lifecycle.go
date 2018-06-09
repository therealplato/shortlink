package lifecycle

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// Cycle provides a context that your application can use at top level.
// The context is terminated when the process receives SIGINT or SIGTERM.
// It also includes a server for kube health checks, off by default
type Cycle struct {
	Ctx context.Context
	End context.CancelFunc
}

// Begin constructs a lifecycle struct and spawns a goroutine to listen for termination signals from OS.
// When a termination signal is received, the lifecycle object calls lifecycle.End(), closing the context.
// It's safe for your code to signal shutdown to other code by calling lifecycle.End() yourself.
// Block on <-lifecycle.Ctx.Done() in your main code path.
// After that unblocks due to the context's channel being closed, you may want to wait longer for slow application code to finish.
func Begin() *Cycle {
	c := &Cycle{}
	c.Ctx, c.End = context.WithCancel(context.Background())
	go c.interceptSignals()
	return c
}

// ServeHTTP returns 200
func (c *Cycle) ServeHTTP(http.ResponseWriter, *http.Request) {
	return
}

// HealthCheck starts listening on listenAddr and responding 200.
// It shuts down within one second of the lifecycle ending.
func (c *Cycle) HealthCheck(listenAddr string) {
	s := http.Server{
		Addr:    listenAddr,
		Handler: c,
	}
	go func() {
		err := s.ListenAndServe()
		if err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()
	<-c.Ctx.Done()
	ctx2, _ := context.WithTimeout(context.Background(), time.Second)
	s.Shutdown(ctx2)
}

func (c *Cycle) interceptSignals() {
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case sig := <-ch:
			log.Printf("lifecycle received %q, terminating\n", sig)
			c.End()
			return
		case <-c.Ctx.Done():
			return
		}
	}

}
