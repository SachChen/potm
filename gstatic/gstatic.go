package gstatic

import (
	"sync"
)

type disk struct {
	Path    string
	Total   int64
	Used    float64
	UsePerc float64
}

type device struct {
	Name    string
	Total   int64
	Used    float64
	UsePerc float64
}

type core struct {
	Class  string
	Number int
}

type Gstatic struct {
	//Addr      string
	Port     []string
	Status   bool
	Hostname string
	Cpu      core     //CPU信息（CPU）
	Cuse     float32  //CPU占用
	Mem      string   //内存大小
	Muse     float32  //内存使用量
	Driver   []device //硬盘信息（硬盘名、总大小以及使用量）
	Mnt      []disk   //挂载信息（挂载名、总大小以及使用量）
}

type Status struct {
	Ginf     map[string]*Gstatic
	GinfLock sync.RWMutex
}

var S = &Status{
	Ginf: make(map[string]*Gstatic),
}

// Aceept 接受guard客户端的信息注册
func Aceept(ip, port, hostname string) {
	//如果guard客户端信息已存在，则修改信息
	if _, ok := S.Ginf[ip]; ok {
		S.GinfLock.Lock()
		S.Ginf[ip].Status = true
		S.Ginf[ip].Hostname = hostname
		//S.Ginf[ip].Port = append(S.Ginf[ip].Port, gjson.Get(info, "Port").Index)
		for _, v := range S.Ginf[ip].Port {
			if v == port {
				S.GinfLock.Unlock()
				return
			}
		}
		S.Ginf[ip].Port = append(S.Ginf[ip].Port, port)
		S.GinfLock.Unlock()

	} else {
		//如果guard客户端信息不存在，则初始化新的客户端信息
		S.GinfLock.Lock()
		S.Ginf[ip] = &Gstatic{Status: true}
		S.Ginf[ip].Hostname = hostname
		//S.Ginf[ip].Port = append(S.Ginf[ip].Port, gjson.Get(info, "Port").Index)
		S.Ginf[ip].Port = append(S.Ginf[ip].Port, port)

		S.GinfLock.Unlock()
	}
	//S.Ginf[ip].Cpu.Class = gjson.Get(info, "Cpu.Class").Str
	//S.Ginf[ip].Cpu.Number = gjson.Get(info, "Cpu.Number").Index
	//S.Ginf[ip].Port = append(S.Ginf[ip].Port, gjson.Get(info, "Port").Index)

}
