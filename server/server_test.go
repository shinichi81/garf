package server_test

import (
	"net/http"
	"testing"

	_ "net/http/pprof"

	"github.com/backenderia/garf/registry"
	"github.com/backenderia/garf/server"
	_echo "github.com/backenderia/garf/server/echo"
	"github.com/labstack/echo"
)

type mockResponseWriter struct{}

func (m *mockResponseWriter) Header() (h http.Header) {
	return http.Header{}
}

func (m *mockResponseWriter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (m *mockResponseWriter) WriteString(s string) (n int, err error) {
	return len(s), nil
}

func (m *mockResponseWriter) WriteHeader(int) {}

func benchRequest(b *testing.B, router http.Handler, r *http.Request) {
	w := new(mockResponseWriter)
	u := r.URL
	rq := u.RawQuery
	r.RequestURI = u.RequestURI()

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		u.RawQuery = rq
		router.ServeHTTP(w, r)
	}
}

func BenchmarkEcho(b *testing.B) {
	reg := registry.New(_echo.New())
	s := reg.Server()

	router := s.Server().(*echo.Echo).Router()

	s.Get("/echo", func(c server.Context) error {
		return c.NoContent(200)
	})

	req, _ := http.NewRequest("GET", "/echo", nil)
	benchRequest(b, router, req)
}
