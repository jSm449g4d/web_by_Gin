package main

import (
	"flag"
	"runtime"

	"github.com/gin-gonic/gin"
)

var restdict = make(map[string]func(c *gin.Context))

func init() {
}

func main() {
	ipport := flag.String("IP_port", "Localhost:20080", "IP:port")
	flag.Parse()

	//r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Static("/static/", "./static/")
	r.LoadHTMLGlob("./templates/*.html")
	r.GET("/", func(c *gin.Context) {
		//c.String(200, "Hello,World!")
		c.HTML(200, "Gin_index.html", gin.H{
			"STATUS_GOLANG_VERSION": runtime.Version(),
			"STATUS_GIN_VERSION":    gin.Version,
		})
	})
	r.GET("/go/:name", func(c *gin.Context) {
		filename := c.Param("name")
		theshow, ok := restdict[filename]
		if ok {
			theshow(c)
		}
	})

	r.GET("/html/:htmlname", func(c *gin.Context) {
		htmlname := c.Param("htmlname")
		print(htmlname)
		c.HTML(200,htmlname+".html" , gin.H{})
	})

	r.Run(*ipport)

}
