package main

import (
	"flag"
	"runtime"

	"github.com/gin-gonic/gin"
)

var restdict = make(map[string]func(c *gin.Context))

func init() {
}

func rest(c *gin.Context) {
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
			"used_golang": runtime.Version(),
			"used_gin":    gin.Version,
		})
	})
	r.GET("/go/:name", rest)

	r.Run(*ipport)

}
