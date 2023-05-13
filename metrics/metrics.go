package metrics

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Metrics(port uint, path string) {
	// 默认情况下 pprof 是不追踪block和mutex的信息的，如果想要看这两个信息，需要手动开启
	runtime.SetBlockProfileRate(1)     // 开启对阻塞操作的跟踪，block
	runtime.SetMutexProfileFraction(1) // 开启对锁调用的跟踪，mutex

	go func() {
		http.Handle(path, promhttp.Handler())
		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			panic(err)
		}
	}()
}
