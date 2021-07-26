package http_demo

import (
    "github.com/matryer/way"
    "net/http"
    "testing"
)

type MyService struct {

}

func (receiver *MyService) ServeHTTP(rw http.ResponseWriter, r *http.Request)  {
    _, _ = rw.Write([]byte("hello world"))
}

func TestHttpWay(t *testing.T) {

    s := MyService{}
    r := way.NewRouter()

    r.Handle("GET", "/get", &s)
    // 监听8080端口
    _ = http.ListenAndServe(":8080", r)

}
