package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)
//基于 errgroup 实现一个 http server 的启动和关闭 ，
//以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出
func main() {
	eg, ctx := errgroup.WithContext(context.Background())

	// 启动 http server
	eg.Go(func() error {
		return Server(ctx, ":8080", NewWorkHandler())
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// signal
	eg.Go(func() error {
		return Signal(ctx, c)
	})

	if err := eg.Wait(); err != nil {
		fmt.Printf("err: %s \n", err.Error())
	}
	fmt.Printf("End")
}

func Server(ctx context.Context, addr string, handler http.Handler) error {
	s := http.Server{
		Addr:    addr,
		Handler: handler,
	}
	go func() {
		<-ctx.Done()
		fmt.Println("Server shutdown")
		s.Shutdown(ctx)
	}()
	return s.ListenAndServe()
}

func Signal(ctx context.Context, c chan os.Signal) error {
	for {
		select {
		case s := <-c:
			return fmt.Errorf("Sys singal: %v \n", s)
		case <-ctx.Done():
			return fmt.Errorf("server quit \n")
		default:
		}
	}
}

type WorkHandler struct {}

func (work *WorkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Server start \n")
}

func NewWorkHandler() *WorkHandler {
	return &WorkHandler{}
}