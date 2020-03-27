package main

import (
	"flag"
	"runtime"

	"github.com/gin-gonic/gin"
)

var restdict = make(map[string]func(c *gin.Context))

func init() {
}

func rests(c *gin.Context) {
	filename := c.Param("name")
	theshow, ok := restdict[filename]
	if ok {
		theshow(c)
	}
}

func main() {
	ipport := flag.String("ipport", "Localhost:20080", "IP:port")
	flag.Parse()

	//r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Static("/static/", "./static/")
	r.LoadHTMLGlob("./templates/*.html")
	r.GET("/", func(c *gin.Context) {
		//c.String(200, "Hello,World!")
		c.HTML(200, "index.html", gin.H{
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

	r.GET("/html/:name", func(c *gin.Context) {
		filename := c.Param("name")
		c.HTML(200, filename+".html", gin.H{})
	})

	r.Run(*ipport)

}
