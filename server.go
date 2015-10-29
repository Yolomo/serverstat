package main

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
)

func base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64Decode(str string) string {
	data, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		fmt.Println("Shit didn't work as planned" + err.Error())
	}
	return string(data)
}

func v1ServerStatUpdate(c *gin.Context) {

	hostname := base64Decode(c.PostForm("hostname"))
	uptime := base64Decode(c.PostForm("uptime"))
	sessions := base64Decode(c.PostForm("sessions"))
	processes := base64Decode(c.PostForm("processes"))
	processes_array := base64Decode(c.PostForm("processes_array"))
	file_handles := base64Decode(c.PostForm("file_handles"))
	file_handles_limit := base64Decode(c.PostForm("file_handles_limit"))
	os_kernel := base64Decode(c.PostForm("os_kernel"))
	os_name := base64Decode(c.PostForm("os_name"))
	os_arch := base64Decode(c.PostForm("os_arch"))
	cpu_name := base64Decode(c.PostForm("cpu_name"))
	cpu_cores := base64Decode(c.PostForm("cpu_cores"))
	cpu_freq := base64Decode(c.PostForm("cpu_freq"))
	ram_total := base64Decode(c.PostForm("ram_total"))
	ram_usage := base64Decode(c.PostForm("ram_usage"))
	swap_total := base64Decode(c.PostForm("swap_total"))
	swap_usage := base64Decode(c.PostForm("swap_usage"))
	disk_array := base64Decode(c.PostForm("disk_array"))
	disk_total := base64Decode(c.PostForm("disk_total"))
	disk_usage := base64Decode(c.PostForm("disk_usage"))
	connections := base64Decode(c.PostForm("connections"))
	nic := base64Decode(c.PostForm("nic"))
	ipv4 := base64Decode(c.PostForm("ipv4"))
	ipv6 := base64Decode(c.PostForm("ipv6"))
	rx := base64Decode(c.PostForm("rx"))
	tx := base64Decode(c.PostForm("tx"))
	rx_gap := base64Decode(c.PostForm("rx_gap"))
	tx_gap := base64Decode(c.PostForm("tx_gap"))
	load := base64Decode(c.PostForm("load"))
	load_cpu := base64Decode(c.PostForm("load_cpu"))
	load_io := base64Decode(c.PostForm("load_io"))
	ping_1 := base64Decode(c.PostForm("ping_1"))

	c.JSON(200, gin.H{
		"status":             "received",
		"hostname":           hostname,
		"uptime":             uptime,
		"sessions":           sessions,
		"processes":          processes,
		"processes_array":    processes_array,
		"file_handles":       file_handles,
		"file_handles_limit": file_handles_limit,
		"os_kernel":          os_kernel,
		"os_name":            os_name,
		"os_arch":            os_arch,
		"cpu_name":           cpu_name,
		"cpu_cores":          cpu_cores,
		"cpu_freq":           cpu_freq,
		"ram_total":          ram_total,
		"ram_usage":          ram_usage,
		"swap_total":         swap_total,
		"swap_usage":         swap_usage,
		"disk_array":         disk_array,
		"disk_total":         disk_total,
		"disk_usage":         disk_usage,
		"connections":        connections,
		"nic":                nic,
		"ipv4":               ipv4,
		"ipv6":               ipv6,
		"rx":                 rx,
		"tx":                 tx,
		"rx_gap":             rx_gap,
		"tx_gap":             tx_gap,
		"load":               load,
		"load_cpu":           load_cpu,
		"load_io":            load_io,
		"ping_1":             ping_1,
	})

}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/serverstat/update", v1ServerStatUpdate)
		v1.GET("/serverstat/ping", func(c *gin.Context) {
			c.String(200, "Pong! ")
		})
	}

	fmt.Println("http://localhost:8120/")
	r.Run(":8120") // listen and serve on 0.0.0.0:8080
}
