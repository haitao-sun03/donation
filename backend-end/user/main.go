package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/haitao-sun03/donation/backend-end/user/config"
	"github.com/haitao-sun03/donation/backend-end/user/routers"
)

func main() {
	config.Init()

	config.RoutinePool.Submit(func() {
		if err := http.ListenAndServe("localhost:6061", nil); err != nil {
			panic(err)
		}
	})

	r := routers.Router()
	config.RoutinePool.Submit(func() {
		if err := r.Run(":10000"); err != nil {
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
