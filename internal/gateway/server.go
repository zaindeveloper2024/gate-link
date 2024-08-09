package gateway

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/zaindeveloper2024/gate-link/cmd/flags"
	"github.com/zaindeveloper2024/gate-link/internal/conf"
	"github.com/zaindeveloper2024/gate-link/internal/middleware"
)

func Handle() {
	http.HandleFunc("/", middleware.Authenticate(middleware.RateLimit(handleProxy)))
	fmt.Printf("API Gateway is running on port %d\n", flags.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", flags.Port), nil))
}

func handleProxy(w http.ResponseWriter, r *http.Request) {
	for _, route := range conf.Conf.Routes {
		if r.URL.Path == route.Path {
			url, err := url.Parse(route.Target)
			if err != nil {
				log.Panicf("Failed to parse route target URL: %v", err)
				http.Error(w, "Failed to parse route target URL", http.StatusInternalServerError)
				return
			}
			proxy := httputil.NewSingleHostReverseProxy(url)
			proxy.ErrorHandler = handleErrorHandler

			r.URL.Host = url.Host
			r.URL.Scheme = url.Scheme
			r.Header.Set("X-Forwarded-Host", r.Header.Get("Host"))
			r.Host = url.Host
			proxy.ServeHTTP(w, r)

			return
		}
	}

	http.NotFound(w, r)
}

func handleErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Printf("Proxy error: %v", err)
	w.WriteHeader(http.StatusBadGateway)
	w.Write([]byte("Proxy error"))
}
