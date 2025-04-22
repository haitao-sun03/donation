package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haitao-sun03/donation/backend-end/config"
	"github.com/haitao-sun03/donation/backend-end/event"
	"github.com/haitao-sun03/donation/backend-end/routers"
)

func main() {
	event.GetEventHandlers()
	ctx, cancel := context.WithCancel(context.Background())
	config.RoutinePool.Submit(func() {
		event.ListenEvents(ctx)
	})

	config.RoutinePool.Submit(func() {
		if err := http.ListenAndServe("localhost:6060", nil); err != nil {
			panic(err)
		}
	})
	r := routers.Router()
	config.RoutinePool.Submit(func() {
		if err := r.Run(":9999"); err != nil {
			panic(err)
		}
	})

	// 捕获退出信号
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	// 等待信号
	<-sigChan
	fmt.Println("Received signal, shutting down...")
	// 清理资源
	cancel()
	config.RoutinePool.ReleaseTimeout(5 * time.Second)
	config.RedisClient.Close()

}
