package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendMessage(id string, trackNumber string, items Items) {
	url := "https://api.line.me/v2/bot/message/push"
	method := "POST"

	//	body model
	// {
	// 	"to":"userid",
	// 	"message":[
	// 		{
	// 			"type":"text",
	// 			"text":"ข้อความที่จะส่ง"
	// 		}
	// 	]
	// }
	var message []map[string]interface{}
	if len(items) == 0 {
		m := make(map[string]interface{})
		m["type"] = "text"
		text := "ไม่พบหมายเลขที่คุณต้องการ"
		m["text"] = text
		message = append(message, m)
	} else {
		m := make(map[string]interface{})
		m["type"] = "text"                               //ปั้นข้อมูล
		text := fmt.Sprintf("หมายเลข : %s", trackNumber) //ปั้นข้อมูล
		m["text"] = text                                 //ปั้นข้อมูล
		message = append(message, m)
		for _, item := range items {
			m := make(map[string]interface{})                                      //ปั้นข้อมูล
			m["type"] = "text"                                                     //ปั้นข้อมูล
			text := fmt.Sprintf("สถานะ : %s\n", item.StatusDescription)            //ปั้นข้อมูล
			text += fmt.Sprintf("ตำแหน่ง : %s %s\n", item.Location, item.Postcode) //ปั้นข้อมูล
			text += fmt.Sprintf("เวลา : %s\n", item.StatusDate)                    //ปั้นข้อมูล
			m["text"] = text
			message = append(message, m)
		}
	}

	payload := make(map[string]interface{})
	payload["to"] = id            //ปั้นข้อมูล
	payload["messages"] = message //ปั้นข้อมูล

	send, err := json.Marshal(payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(string(send)))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
