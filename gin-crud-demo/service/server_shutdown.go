package service
import (
"context"
"log"
"net/http"
"os"
"os/signal"
"syscall"
)

// 优雅关闭应用程序

func GracefulShutdown() {
	// 在此处执行任何其他清理操作
	log.Println("Shutting down gracefully...")
}


// 优雅关闭应用程序

func insideGracefulShutdown() {
	// 在此处执行任何其他清理操作
	log.Println("Shutting down gracefully...")
}
// 等待应用程序退出信号

func WaitForShutdown(srv *http.Server) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("Received shutdown signal. Shutting down...")

	// 关闭HTTP服务器
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Fatalf("Error shutting down server: %v", err)
	}

	insideGracefulShutdown()
}
