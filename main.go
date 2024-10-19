package main

import "fmt"

func main() {
    type Backend struct {
        URL *url.URL
        Alive bool
        mux sync.RWMutex
        ReverseProxy *httputil.ReverseProxy
    }

    type ServerPool struct {
        backends []*Backend
        current uint64
    }

    u, _ := url.Parse("http://localhost:8080")
    rp := httputil.NewSingleHostReverseProxy(u)

    http.HandlerFunc(rp.ServeHTTP)

    func (s *ServePool) NextIndex() int {
        return int(atomic.AddUint64(&s.current, uint64(1)) % uint64(len(s.backends)))
    }
}
