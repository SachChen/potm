package gstatic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type List struct {
	File   string
	Status bool
	Eid    int
	Rule   string
}

// show 获取guard返回的所有app的信息
func Showcron(ip, port string) string {
	Info := make(map[string]List)
	resp, err := http.Get("http://" + ip + ":" + port + "/seecron")
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
