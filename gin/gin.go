package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//	os.Chdir(exe)
	//r := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.LoadHTMLGlob("./templates/*.html")

	r.GET("/", func(c *gin.Context) {
		//c.String(200, "Hello,World!")
		c.HTML(200, "index.html", gin.H{})
	})

	r.Run("Localhost:20080")

}
