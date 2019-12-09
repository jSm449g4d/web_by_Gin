package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	restdict["22"] = func(c *gin.Context) {
		v, _ := mem.VirtualMemory()
		cpup, _ := cpu.Percent(0, true)
		cpuf, _ := cpu.Info()
		c.HTML(200, "22.html", gin.H{
			"CPU_M": cpuf[0].ModelName,
			"CPU_F": cpuf[0].Mhz,
			"CPU_P": cpup,
			"MEM_T": v.Total,
			"MEM_U": v.Used,
		})
	}
}
