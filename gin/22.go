package main

import (
	"github.com/gin-gonic/gin"
)

func init() {
	restdict["22.go"] = func(c *gin.Context) {
		//v, _ := mem.VirtualMemory()
		//cpup, _ := cpu.Percent(0, true)
		//cpuf, _ := cpu.Info()
		c.HTML(200, "22.html", gin.H{
			//"CPU_M": cpuf[0].ModelName,
			//"CPU_F": cpuf[0].Mhz,
			//"CPU_P": cpup,
			//"MEM_T": v.Total,
			//"MEM_U": v.Used,
		})
	}
}
