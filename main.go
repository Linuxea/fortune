package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
)

func main() {

	b, err2 := exec.Command("fortune").Output()
	if err2 != nil {
		panic(err2)
	}

	content := `
	{
        "msgtype": "text",
        "text": {
        "content": ` + "\"" + string(b) + "\"" + `
        }
   }
   `

	fmt.Println(content)

	r, err := http.Post("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=ea986c1a-4832-41f2-a170-516d028d8641", "application/json", bytes.NewBuffer([]byte(content)))

	if err != nil {
		panic(err)
	}

	rr := make([]byte, 0)
	nn, err := r.Body.Read(rr)
	if err != nil {
		panic(err)
	}

	fmt.Println(nn)

	fmt.Println(string(rr))

	r.Body.Close()

}
