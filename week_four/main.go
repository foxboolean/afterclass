package main

import (
	"afterclass/week_four/api"
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

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

type WorkHandler struct{}

func (work *WorkHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	fmt.Printf("Server start \n")
	ua := api.NewUserApi(InitUserService())
	result := ua.QueryUserInfo(api.NewRequestInfo("1"))
	fmt.Printf("result is %+v", result)
}

func NewWorkHandler() *WorkHandler {
	return &WorkHandler{}
}
