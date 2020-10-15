package module

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func TimeToCreatelogDir() {
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		SendTextMessage()
	}

}
func SendTextMessage() {

	url := "http://11.8.75.19:12345/telephonestatus"
	method := "POST"

	payload := strings.NewReader("{ \n    \"status\": \"1\", \n    \"msg\": \"test\",\n    \"alarm\":{\n        \"endpoint\":\"11.8.75.73\",\n        \"sname\":\"内存占用超过75%\",\n        \"event_type\":\"alert\",\n        \"phone\":\"12345678910\"\n        }\n}")

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
}
