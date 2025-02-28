package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haitao-sun03/go/event"

	"github.com/haitao-sun03/go/config"
	"github.com/haitao-sun03/go/routers"
)

func main() {

	config.Init()
	config.RoutinePool.Submit(func() {
		event.ListenEvents()
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
	config.RoutinePool.ReleaseTimeout(5 * time.Second)
	config.RedisClient.Close()
}
