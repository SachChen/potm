package main

import (
	//"fmt"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mgua/gstatic"
	"net/http"

	"github.com/gin-gonic/gin"
)

type List struct {
	Pid       int
	Status    bool
	Logsize   int
	Logfiles  int
	Alart     string
	Logapi    string
	Logserver string
	Topic     string
	WashMode  string
	Version   string
	Startup   string
	Dure      int
	Trytime   int
}

// show 获取guard返回的所有app的信息
func show(ip, port string) string {
	Info := make(map[string]List)
	resp, err := http.Get("http://" + ip + ":" + port + "/allinfo")
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
	}
	defer resp.Body.Close()

	bodyContent, err := ioutil.ReadAll(resp.Body)

	err1 := json.Unmarshal([]byte(bodyContent), &Info) //第二个参数要地址内存传递
	if err1 != nil {
		fmt.Println("err = ", err1)
	}
	jsoni, _ := json.Marshal(Info)
	stringi := string(jsoni)
	return stringi
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		//放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}

func main() {

	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	r.Use(Cors())

	r.Static("/static", "./static")

	r.LoadHTMLGlob("view/*")

	r.GET("/", func(c *gin.Context) {

		guinfo, _ := json.Marshal(gstatic.S.Ginf)

		c.HTML(http.StatusOK, "index.html", gin.H{
			"ginfo": string(guinfo),
		})

	})

	r.GET("/status", func(c *gin.Context) {
		ip := c.DefaultQuery("ip", "127.0.0.1")
		port := c.DefaultQuery("port", "8010") // shortcut for     c.Request.URL.Query().Get("lastname")
		c.HTML(http.StatusOK, "check.html", gin.H{
			"IP":   ip,
			"PORT": port,
			"Info": show(ip, port),
		})
	})

	r.GET("/cron", func(c *gin.Context) {
		ip := c.DefaultQuery("ip", "127.0.0.1")
		port := c.DefaultQuery("port", "8010") // shortcut for     c.Request.URL.Query().Get("lastname")
		c.HTML(http.StatusOK, "croninfo.html", gin.H{
			"IP":   ip,
			"PORT": port,
			"Info": gstatic.Showcron(ip, port),
		})
	})

	/*
		r.GET("/appinfo", func(c *gin.Context) {
			ip := c.DefaultQuery("ip", "127.0.0.1")
			port := c.DefaultQuery("port", "8010")
			c.HTML(http.StatusOK, "check.html", gin.H{
				"IP":   ip,
				"PORT": port,
				"Info": show(ip, port),
			})
		})
	*/

	//通过/seelog查看日志
	r.GET("/seelog", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		lines := c.DefaultQuery("lines", "1000")
		app := c.Query("app")
		resp, err := http.Get("http://" + ip + ":" + port + "/seelog?file=" + app + "&lines=" + lines)
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			return
		}
		defer resp.Body.Close()
		bodyContent, err := ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK, string(bodyContent))
	})

	//通过/regist注册guard信息
	r.GET("/regist", func(c *gin.Context) {
		ip := c.ClientIP()
		port := c.DefaultQuery("port", "8010")
		hostname := c.Query("name")
		//gport, _ := strconv.Atoi(port)
		gstatic.Aceept(ip, port, hostname)
	})

	// 资源信息页面展示
	r.GET("/hostinfo", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		//gport, _ := strconv.Atoi(port)
		sys := gstatic.Getres(ip, port)
		//fmt.Println(sys)
		jsoni, _ := json.Marshal(sys)
		stringi := string(jsoni)
		c.HTML(http.StatusOK, "resource.html", gin.H{
			"IP":   ip,
			"PORT": port,
			"Info": stringi,
		})
	})

	// 返回硬件资源信息
	r.GET("/getres", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		resp, err := http.Get("http://" + ip + ":" + port + "/getres")
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			return
		}
		defer resp.Body.Close()
		bodyContent, err := ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK, string(bodyContent))
	})

	r.GET("/getsys", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		resp, err := http.Get("http://" + ip + ":" + port + "/getsys")
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			return
		}
		defer resp.Body.Close()
		bodyContent, err := ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK, string(bodyContent))
	})

	r.GET("/getproc", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		app := c.Query("app")
		resp, err := http.Get("http://" + ip + ":" + port + "/getproc?file=" + app)
		//fmt.Println("http://" + ip + ":" + port + "/getproc&file=" + app)
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			return
		}
		defer resp.Body.Close()
		bodyContent, err := ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK, string(bodyContent))
	})

	r.GET("/pstart", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		app := c.Query("app")
		resp, err := http.Get("http://" + ip + ":" + port + "/launch?file=" + app)
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			fmt.Printf("http://" + ip + ":" + port + "/launch?file=" + app)
			return
		}
		defer resp.Body.Close()

	})

	r.GET("/pstop", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		app := c.Query("app")
		resp, err := http.Get("http://" + ip + ":" + port + "/down?file=" + app)
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			return
		}
		defer resp.Body.Close()

	})

	r.GET("/allproc", func(c *gin.Context) {
		ip := c.Query("ip")
		port := c.DefaultQuery("port", "8010")
		resp, err := http.Get("http://" + ip + ":" + port + "/allproc")
		if err != nil {
			fmt.Printf("get request failed, err:[%s]", err.Error())
			return
		}
		defer resp.Body.Close()
		bodyContent, err := ioutil.ReadAll(resp.Body)
		c.String(http.StatusOK, string(bodyContent))

	})

	r.Run()
}
