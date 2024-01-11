{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

// type myHandler {
//     {{$svrType}}GinClientHandlerImpl
//	   Log *slog.Logger	
// }
// 
// func NewMyHandler(log *slog.Logger) *myHandler {
// 	   return &myHandler{Log: log}
// }
// 
// func main () {
//     e := gin.Default()
//     log := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{}))
//     handler := example.NewExampleServiceGinServerHandler(NewMyHandler(log))
//     e.Get("/", e.handler.Hello)
//     if err := e.Start(":1111"); err != nil {
//         panic(err)
//     }
// }
// 
// func (h myHandler) Hello(c *gin.Context, args *example.HelloRequest) (*example.HelloResponse, error) {
//     h.Log.Info("打印参数", slog.Any("args", args))
//     return &example.HelloResponse{Msg: args.Name}, nil
// }

type {{$svrType}}GinServerHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(*gin.Context)
{{- end}}
}

type {{$svrType}}GinServerHandlerImpl struct {
	Handler {{$svrType}}GinClientHandler
}

func New{{$svrType}}GinServerHandler(handler {{$svrType}}GinClientHandler) {{$svrType}}GinServerHandler {
	return &{{$svrType}}GinServerHandlerImpl{
		Handler: handler,
	}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}}
{{- end}}
{{- if ne .Swag ""}}
{{ .Swag }}
{{- end}}
func (s {{$svrType}}GinServerHandlerImpl) {{.Name}}(c *gin.Context) {
	var args {{.Request}}
	if err := c.{{.Bind}}(&args); err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	resp, err := s.Handler.{{.Name}}(c, &args)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": resp})
}
{{- end}}

type {{$svrType}}GinClientHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(*gin.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

// 下面方法仅供参考, 具体需要自己实现

type {{$svrType}}GinClientHandlerImpl struct {
}

func New{{$svrType}}GinClientHandler() {{$svrType}}GinClientHandler {
	return &{{$svrType}}GinClientHandlerImpl{}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}}
{{- end}}
func ({{$svrType}}GinClientHandlerImpl) {{.Name}}(c *gin.Context, args *{{.Request}}) (*{{.Reply}}, error) {
	return &{{.Reply}}{}, nil
}
{{- end}}