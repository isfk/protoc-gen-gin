# isfk/protoc-gen-gin

## 安装

```sh
go install github.com/isfk/protoc-gen-gin@latest
```

## 生成

```sh
protoc --go_out=. --go_opt=paths=source_relative \
    --gin_out=. --gin_opt=paths=source_relative \
    demo/demo.proto

# go build .; cp protoc-gen-gin ~/Projects/bin; buf generate example; buf generate example --template=buf.gen.tag.yaml
```

## 测试代码

```go
package main

import (
	"os"

	"github.com/isfk/protoc-gen-gin/example"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
)

type myHandler struct {
	example.ExampleService_GinClientHandlerImpl
	log *slog.Logger
}

func NewMyHandler(log *slog.Logger) *myHandler {
	return &myHandler{log: log}
}

func main() {
	e := gin.Default()

	log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))

	handler := example.NewExampleService_GinServerHandler(NewMyHandler(log))

	e.GET("/hello", handler.Hello)
	e.GET("/say", handler.Say)
	e.Start(":1111")
}

func (h myHandler) Hello(args *example.HelloRequest) (*example.HelloResponse, error) {
	h.log.Info("打印参数", slog.Any("args", args))
	return &example.HelloResponse{Msg: args.Name}, nil
}

func (h myHandler) Say(args *example.SayRequest) (*example.SayResponse, error) {
	h.log.Info("打印参数", slog.Any("args", args))
	return &example.SayResponse{Msg: args.Name}, nil
}
```

## 调试

使用 `fmt.Fprintf`

```go
fmt.Fprintf(os.Stderr, "%v \n", method.Desc.Name())
fmt.Fprintf(os.Stderr, "%v \n", method.Desc.FullName())
```
