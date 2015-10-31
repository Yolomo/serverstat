package main

import (
	"encoding/base64"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

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

func getJsonVar(js *simplejson.Json, str string) string {
	result, err := js.Get(str).String()
	check(err)
	return result
}

func v1ViewByUID(c *gin.Context) {
	uid := c.Param("uid")
	dat, err := ioutil.ReadFile("hosts/host-" + uid + ".json")
	check(err)

	fmt.Print(string(dat))

	js, err := simplejson.NewJson(dat)
	check(err)

	hostname := getJsonVar(js, "hostname")
	uptime := getJsonVar(js, "uptime")
	ipv4 := getJsonVar(js, "ipv4")

	c.JSON(200, gin.H{
		"hostname": hostname,
		"uptime":   uptime,
		"ipv4":     ipv4,
		"osname":   getJsonVar(js, "osname"),
	})

}
func v1ServerStatUpdate(c *gin.Context) {

	uid := base64Decode(c.PostForm("uid"))
	hostname := base64Decode(c.PostForm("hostname"))
	uptime := base64Decode(c.PostForm("uptime"))
	sessions := base64Decode(c.PostForm("sessions"))
	processes := base64Decode(c.PostForm("processes"))
	processesarray := base64Decode(c.PostForm("processesarray"))
	filehandles := base64Decode(c.PostForm("filehandles"))
	filehandleslimit := base64Decode(c.PostForm("filehandleslimit"))
	oskernel := base64Decode(c.PostForm("oskernel"))
	osname := base64Decode(c.PostForm("osname"))
	osarch := base64Decode(c.PostForm("osarch"))
	cpuname := base64Decode(c.PostForm("cpuname"))
	cpucores := base64Decode(c.PostForm("cpucores"))
	cpufreq := base64Decode(c.PostForm("cpufreq"))
	ramtotal := base64Decode(c.PostForm("ramtotal"))
	ramusage := base64Decode(c.PostForm("ramusage"))
	swaptotal := base64Decode(c.PostForm("swaptotal"))
	swapusage := base64Decode(c.PostForm("swapusage"))
	diskarray := base64Decode(c.PostForm("diskarray"))
	disktotal := base64Decode(c.PostForm("disktotal"))
	diskusage := base64Decode(c.PostForm("diskusage"))
	connections := base64Decode(c.PostForm("connections"))
	nic := base64Decode(c.PostForm("nic"))
	ipv4 := base64Decode(c.PostForm("ipv4"))
	ipv6 := base64Decode(c.PostForm("ipv6"))
	rx := base64Decode(c.PostForm("rx"))
	tx := base64Decode(c.PostForm("tx"))
	rxgap := base64Decode(c.PostForm("rxgap"))
	txgap := base64Decode(c.PostForm("txgap"))
	load := base64Decode(c.PostForm("load"))
	loadcpu := base64Decode(c.PostForm("loadcpu"))
	loadio := base64Decode(c.PostForm("loadio"))
	ping1 := base64Decode(c.PostForm("ping1"))

	c.JSON(200, gin.H{
		"status":           "received",
		"uid":              uid,
		"hostname":         hostname,
		"uptime":           uptime,
		"sessions":         sessions,
		"processes":        processes,
		"processesarray":   processesarray,
		"filehandles":      filehandles,
		"filehandleslimit": filehandleslimit,
		"oskernel":         oskernel,
		"osname":           osname,
		"osarch":           osarch,
		"cpuname":          cpuname,
		"cpucores":         cpucores,
		"cpufreq":          cpufreq,
		"ramtotal":         ramtotal,
		"ramusage":         ramusage,
		"swaptotal":        swaptotal,
		"swapusage":        swapusage,
		"diskarray":        diskarray,
		"disktotal":        disktotal,
		"diskusage":        diskusage,
		"connections":      connections,
		"nic":              nic,
		"ipv4":             ipv4,
		"ipv6":             ipv6,
		"rx":               rx,
		"tx":               tx,
		"rxgap":            rxgap,
		"txgap":            txgap,
		"load":             load,
		"loadcpu":          loadcpu,
		"loadio":           loadio,
		"ping1":            ping1,
	})

}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{
			"title":    "Serverstat",
			"subtitle": "Easy server statistics monitoring",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.POST("/serverstat/update", v1ServerStatUpdate)
		v1.GET("/serverstat/view/:uid/", v1ViewByUID)
		v1.GET("/serverstat/ping", func(c *gin.Context) {
			c.String(200, "Pong! ")
		})
	}

	fmt.Println("http://localhost:8120/")
	r.Run(":8120") // listen and serve on 0.0.0.0:8080
}
