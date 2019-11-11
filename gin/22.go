package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

func init() {
	restdict["22"] = func(c *gin.Context) {
		v, _ := mem.VirtualMemory()
		cpus, _ := cpu.Percent(0, true)
		c.HTML(200, "22.html", gin.H{
			"CPU_P": cpus,
			"MEM_T": v.Total,
			"MEM_U": v.Used,
		})
	}
}
