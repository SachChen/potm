package gstatic

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Sys struct {
	Hostname string
	Os       string
	Release  string
	Arch     string
	Systime  string
	Zone     string
}

func Getres(ip, port string) map[string]string {
	Info := make(map[string]string)
	resp, err := http.Get("http://" + ip + ":" + port + "/getsys")
	if err != nil {
		fmt.Printf("get request failed, err:[%s]", err.Error())
	}
	defer resp.Body.Close()
	bodyContent, err := ioutil.ReadAll(resp.Body)

	err1 := json.Unmarshal([]byte(bodyContent), &Info)
	if err1 != nil {
		fmt.Println("err = ", err1)
	}

	return Info
	/*
		jsoni, _ := json.Marshal(Info)
		stringi := string(jsoni)
		fmt.Println(stringi)
		return stringi
	*/
}
