package goauditlogger

import (
	"encoding/json"
	"io"
	"os"
	"time"

	"github.com/gofiber/fiber"
)

type Config struct {
	Next             func(c *fiber.Ctx) bool
	Done             func(c *fiber.Ctx, logString []byte)
	Format           string
	TimeFormat       string
	TimeZone         string
	TimeInterval     time.Duration
	Output           io.Writer
	enableColors     bool
	enableLatency    bool
	timeZoneLocation *time.Location
}

type FiberLoggerMiddleWare struct {
	TagPid               string `json:"pid"`
	TagTime              string `json:"time"`
	TagReferer           string `json:"referer"`
	TagProtocol          string `json:"protocol"`
	TagPort              string `json:"port"`
	TagIP                string `json:"ip"`
	TagIPs               string `json:"ips"`
	TagHost              string `json:"host"`
	TagMethod            string `json:"method"`
	TagPath              string `json:"path"`
	TagURL               string `json:"url"`
	TagUA                string `json:"ua"`
	TagLatency           string `json:"latency"`
	TagStatus            string `json:"status"`  // response status
	TagResBody           string `json:"resBody"` // response body
	TagReqHeaders        string `json:"reqHeaders"`
	TagQueryStringParams string `json:"queryParams"` // request query parameters
	TagBody              string `json:"body"`        // request body
	TagBytesSent         string `json:"bytesSent"`
	TagBytesReceived     string `json:"bytesReceived"`
	TagRoute             string `json:"route"`
	TagError             string `json:"error"`
}

func Get_DefaultFormat() string {
	config_format := FiberLoggerMiddleWare{
		TagPid:               "${pid}",
		TagTime:              "${time}",
		TagReferer:           "${referer}",
		TagProtocol:          "${protocol}",
		TagPort:              "${port}",
		TagIP:                "${ip}",
		TagIPs:               "${ips}",
		TagHost:              "${host}",
		TagMethod:            "${method}",
		TagPath:              "${path}",
		TagURL:               "${url}",
		TagUA:                "${ua}",
		TagLatency:           "${latency}",
		TagStatus:            "${status}",
		TagResBody:           "${resBody}",
		TagReqHeaders:        "${reqHeaders}",
		TagQueryStringParams: "${queryParams}",
		TagBody:              "${body}",
		TagBytesSent:         "${bytesSent}",
		TagBytesReceived:     "${bytesReceived}",
		TagRoute:             "${route}",
		TagError:             "${error}",
	}
	logStr, _ := json.Marshal(config_format)
	return string(logStr) + "\n"
}

func Get_Output() *os.File {
	File, _ := os.OpenFile("./audit.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	return File
}

func Get_DefaultConfig() Config {
	var config = Config{
		Next:         nil,
		Done:         nil,
		Format:       Get_DefaultFormat(),
		TimeFormat:   "15:04:05",
		TimeZone:     "Local",
		TimeInterval: 500 * time.Millisecond,
		Output:       Get_Output(),
	}
	return config
}
