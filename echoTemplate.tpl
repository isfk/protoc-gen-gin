{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}

// type myHandler {
//     {{$svrType}}_GinClientHandlerImpl
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
//     handler := example.NewExampleService_GinServerHandler(NewMyHandler(log))
//     e.Get("/", e.handler.Hello)
//     if err := e.Start(":1111"); err != nil {
//         panic(err)
//     }
// }
// 
// func (h myHandler) Hello(args *example.HelloRequest) (*example.HelloResponse, error) {
//     h.Log.Info("打印参数", slog.Any("args", args))
//     return &example.HelloResponse{Msg: args.Name}, nil
// }

type {{$svrType}}_GinServerHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(*gin.Context)
{{- end}}
}

type {{$svrType}}_GinServerHandlerImpl struct {
	Handler {{$svrType}}_GinClientHandler
}

func New{{$svrType}}_GinServerHandler(handler {{$svrType}}_GinClientHandler) {{$svrType}}_GinServerHandler {
	return &{{$svrType}}_GinServerHandlerImpl{
		Handler: handler,
	}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}}
{{- end}}
{{ .Swag }}
func (s {{$svrType}}_GinServerHandlerImpl) {{.Name}}(c *gin.Context) {
	var args {{.Request}}
	if err := c.Bind(&args); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	resp, err := s.Handler.{{.Name}}(&args)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 500, "msg": err.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": resp})
}
{{- end}}

type {{$svrType}}_GinClientHandler interface {
{{- range .Methods}}
	{{- if ne .Comment ""}}
	{{.Comment}}
	{{- end}}
	{{.Name}}(*{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

// 下面方法仅供参考, 具体需要自己实现

type {{$svrType}}_GinClientHandlerImpl struct {
}

func New{{$svrType}}_GinClientHandler() {{$svrType}}_GinClientHandler {
	return &{{$svrType}}_GinClientHandlerImpl{}
}

{{- range .Methods}}
{{- if ne .Comment ""}}
{{.Comment}}
{{- end}}
func ({{$svrType}}_GinClientHandlerImpl) {{.Name}}(args *{{.Request}}) (*{{.Reply}}, error) {
	return &{{.Reply}}{}, nil
}
{{- end}}