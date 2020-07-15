package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/khanamoto/dokodemo/config"
	"github.com/khanamoto/dokodemo/repository"
	"github.com/khanamoto/dokodemo/service"
	"github.com/khanamoto/dokodemo/web"
)

func main() {
	if err := run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

// error型
func run(_ []string) error {
	conf, err := config.Load()
	if err != nil {
		return fmt.Errorf("failed to load config: %+v", err)
	}

	repo, err := repository.New(conf.DbDsn)
	if err != nil {
		return fmt.Errorf("failed to create repository: %+v", err)
	}

	app := service.NewApp(repo)
	server := &http.Server{
		Addr:    ":" + strconv.Itoa(conf.Port),
		Handler: web.NewServer(app).Handler(),
	}

	fmt.Printf("Starting server...(config: %#v)\n", conf)
	// goルーチン
	go graceful(server, 10*time.Second)
	if err = server.ListenAndServe(); err != http.ErrServerClosed {
		return err
	}
	return nil
}

// 正常終了(Graceful Shutdown)
func graceful(server *http.Server, timeout time.Duration) {
	// チャネル
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	sig := <-sigChan
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	fmt.Println("shutting down server...", sig)
	if err := server.Shutdown(ctx); err != nil {
		fmt.Printf("failed to shutdown: %v\n", err)
	}
}