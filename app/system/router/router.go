package router

import (
	"github.com/cilidm/toolbox/logging"
	"github.com/gin-gonic/gin"
	"oss_manage/app/system/pkg"
	"oss_manage/app/system/pkg/config"
	"oss_manage/app/system/pkg/session"
)

//func InitSession() gin.HandlerFunc {
//	//store := memstore.NewStore([]byte(session.DefaultKey))	// 使用内存保存session
//	//store := cookie.NewStore([]byte(session.DefaultKey))  		// 使用cookie保存session
//	store,err := redis.NewStore(100,"tcp","127.0.0.1:6379","",[]byte(session.DefaultKey))
//	if err != nil{
//		logging.Fatal(err.Error())
//	}
//	return session.Sessions("_SESSION",store)
//}

var (
	GroupList     = make([]*routerGroup, 0)
	PermissionMap = make(map[string]string, 0)
	HtmlTemplate  = make([]string, 0)
)

func InitRouter() *gin.Engine {
	gin.SetMode(config.Setting.RunMode)
	r := gin.New()
	// 图片上传文件存储
	r.Static(config.Setting.ImgUrlPath, config.Setting.ImgSavePath)

	r.Static("/static", "static")

	//r.LoadHTMLGlob("view/**/**/*")
	InitTemplate("view/system")
	if len(HtmlTemplate) > 0 { // 引用其他模块前端页面
		r.LoadHTMLFiles(HtmlTemplate...)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//r.Use(middleware.Cors())  //暂不配置cors
	//r.Use(InitSession())
	r.Use(session.EnableCookieSession(config.Setting.JwtSecret))

	if len(GroupList) > 0 { // 通过 _ 引入system/controller下的init router
		for _, group := range GroupList {
			g := r.Group(group.UlrPath, group.Handlers...)
			for _, r2 := range group.Router {
				g.Handle(r2.Method, r2.UrlPath, r2.HandlerFunc...)
			}
		}
	}
	return r
}

func InitTemplate(temPath string) {
	files, err := pkg.GetTemplateFiles(temPath)
	if err != nil {
		logging.Fatal("未找到前端模板文件:" + err.Error())
	}
	for _, v := range files {
		HtmlTemplate = append(HtmlTemplate, v)
	}
}

const (
	GET     = "GET"
	POST    = "POST"
	PUT     = "PUT"
	PATCH   = "PATCH"
	HEAD    = "HEAD"
	OPTIONS = "OPTIONS"
	DELETE  = "DELETE"
	CONNECT = "CONNECT"
	TRACE   = "TRACE"
)

type router struct {
	Method      string
	UrlPath     string
	HandlerFunc []gin.HandlerFunc
}

type routerGroup struct {
	ServerName string            //服务名称
	UlrPath    string            //URL路径
	Handlers   []gin.HandlerFunc //中间件
	Router     []*router         //路由
}

func New(serverName, urlPath string, middleware ...gin.HandlerFunc) *routerGroup {
	var r routerGroup
	r.ServerName = serverName
	r.Router = make([]*router, 0)
	r.UlrPath = urlPath
	r.Handlers = middleware
	GroupList = append(GroupList, &r)
	return &r
}

func (group *routerGroup) Handle(method, urlPath string, handlers ...gin.HandlerFunc) *routerGroup {
	var r router
	r.Method = method
	r.UrlPath = urlPath
	r.HandlerFunc = handlers
	group.Router = append(group.Router, &r)
	return group
}

//添加路由信息-ANY
func (group *routerGroup) ANY(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle("ANY", relativePath, handlers...)
	return group
}

//添加路由信息-GET
func (group *routerGroup) GET(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(GET, relativePath, handlers...)
	return group
}

//添加路由信息-POST
func (group *routerGroup) POST(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(POST, relativePath, handlers...)
	return group
}

//添加路由信息-OPTIONS
func (group *routerGroup) OPTIONS(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(OPTIONS, relativePath, handlers...)
	return group
}

//添加路由信息-PUT
func (group *routerGroup) PUT(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PUT, relativePath, handlers...)
	return group
}

//添加路由信息-PATCH
func (group *routerGroup) PATCH(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(PATCH, relativePath, handlers...)
	return group
}

//添加路由信息-HEAD
func (group *routerGroup) HEAD(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(HEAD, relativePath, handlers...)
	return group
}

//添加路由信息-DELETE
func (group *routerGroup) DELETE(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(DELETE, relativePath, handlers...)
	return group
}

//添加路由信息-CONNECT
func (group *routerGroup) CONNECT(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(CONNECT, relativePath, handlers...)
	return group
}

//添加路由信息-TRACE
func (group *routerGroup) TRACE(relativePath string, handlers ...gin.HandlerFunc) *routerGroup {
	group.Handle(TRACE, relativePath, handlers...)
	return group
}
