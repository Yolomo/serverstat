package dataprocessor

import (
/*	"database/sql"
	"encoding/base64"
	"encoding/json"
	"fmt"*/
)

type host struct {
	UID              string
	Connections      string `json:"connections"`
	Cpucores         string `json:"cpucores"`
	Cpufreq          string `json:"cpufreq"`
	Cpuname          string `json:"cpuname"`
	Diskarray        string `json:"diskarray"`
	Disktotal        string `json:"disktotal"`
	Diskusage        string `json:"diskusage"`
	Filehandles      string `json:"filehandles"`
	Filehandleslimit string `json:"filehandleslimit"`
	Hostname         string `json:"hostname"`
	Ipv4             string `json:"ipv4"`
	Ipv6             string `json:"ipv6"`
	Lastupdate       string `json:"lastupdate"`
	Load             string `json:"load"`
	Loadcpu          string `json:"loadcpu"`
	Loadio           string `json:"loadio"`
	Nic              string `json:"nic"`
	Osarch           string `json:"osarch"`
	Oskernel         string `json:"oskernel"`
	Osname           string `json:"osname"`
	Ping1            string `json:"ping1"`
	Processes        string `json:"processes"`
	Processesarray   string `json:"processesarray"`
	Ramtotal         string `json:"ramtotal"`
	Ramusage         string `json:"ramusage"`
	Rx               string `json:"rx"`
	Rxgap            string `json:"rxgap"`
	Sessions         string `json:"sessions"`
	Swaptotal        string `json:"swaptotal"`
	Swapusage        string `json:"swapusage"`
	Tx               string `json:"tx"`
	Txgap            string `json:"txgap"`
	Uptime           string `json:"uptime"`
}
