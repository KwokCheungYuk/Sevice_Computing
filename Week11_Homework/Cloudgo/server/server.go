package server

import(
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/martini-contrib/binding"
)

type Post struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

func Begin(port string) { 

	//默认设置
	defaultConfig := martini.Classic()

	
	//设置资源文件的路径
	defaultConfig.Use(martini.Static("assets"))
	//设置渲染html模板的包
	defaultConfig.Use(render.Renderer())
	
	//GET请求
	defaultConfig.Get("/", func(r render.Render) {
		//获取当前系统时间
		curTime:=time.Now().Format("2006-01-02 15:04:05")
		//渲染index模板，将时间传过去
        r.HTML(200, "index", map[string]interface{}{"Time":curTime})
	})

	//POST请求
	defaultConfig.Post("/", binding.Bind(Post{}), func(post Post, r render.Render) {
		curTime:=time.Now().Format("2006-01-02 15:04:05")
		p := Post{Username: post.Username, Password: post.Password}
		//渲染index模板，将用户信息和时间传过去
        r.HTML(200, "info", map[string]interface{}{"post": p, "Time":curTime})
	})

	//开始运行
	defaultConfig.RunOnAddr(":"+port)
}