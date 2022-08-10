package decorator

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// 实现http中间件记录请求的URL/方法
// 实现http中间件记录请求的网路地址
// 实现http中间件记录请求的耗时

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello decorator")
}

func logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("logging start")
		log.Printf("记录请求的URL和方法： %s, %s", r.URL, r.Method)
		next.ServeHTTP(w, r)
		log.Printf("logging end")
	})
}

func tracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		next.ServeHTTP(w, r)
		elapsedTime := time.Since(startTime)
		log.Printf("记录方法的执行时间 %s", elapsedTime)
	})
}

func main() {
	http.Handle("/", logging(tracing(http.HandlerFunc(hello))))

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Printf("start server err: %s", err)
		return
	}
}
