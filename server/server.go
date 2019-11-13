package server

import (
	"time"
	"github.com/go-martini/martini"
	"github.com/martini-contrib/binding"
	"github.com/martini-contrib/render"
)

type Post struct {
	MyWord string `form:"MyWord" binding:"required"`
}//用binding包定义一个结构体Post

func Begin(port string) {
	m := martini.Classic() //创建一个典型的Martini实例m

	m.Use(martini.Static("assets")) //设置图片路径
	m.Use(render.Renderer())        //设置渲染html模板的包

	m.Get("/", func(r render.Render) {
		t := time.Now().Format(time.UnixDate)
		r.HTML(200, "index", map[string]interface{}{"Time": t})
	}) //GET请求，获取当前时间，并渲染index模板，传入获取的时间信息

	m.Post("/", binding.Bind(Post{}), func(post Post, r render.Render) {
		p := Post{MyWord: post.MyWord}
		r.HTML(200, "info", map[string]interface{}{"post": p})
	}) //POST请求，渲染info模板，传入页面1填写的话语内容

	m.RunOnAddr(":" + port) //在端口开始运行
}
