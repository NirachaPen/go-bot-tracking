package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// function Tracking  เอาไว้ ยิง api ไปดึงข้อมูล จาก ไปรศีไทย โดยวิธีการ http request โดยมี พารามิเจอร์
// 1 หมายเลข tracking  (trackNumber)
func Tracking(tracking string) TrackData {
	// url ในการ ยิง request เพื่อไปดึงข้อมูลจาก ไปรศณี
	url := "https://trackapi.thailandpost.co.th/post/api/v1/track"
	// method ในการ ยิง request เพื่อไปดึงข้อมูลจาก ไปรศณี
	method := "POST"
	//รูปแบบข้อมูลที่จะส่ง ดึงข้อมูล ของ ไปรณี
	//api ไปรณี กำหนดมา
	payload := strings.NewReader(`{
		"status": "all",
		"language": "TH",
		"barcode": [
			"` + tracking + `"
		]
	}`)
	var track TrackData      // ประการตัวแปร ที่มี struct เป้ฯ TrackData ปั้น struct ตาม response ที่ api ส่งมาไห้
	client := &http.Client{} // postman generate มาไห้ เป็น การ ยิง request api
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return track
	}
	// ใส่ token เข้าไปด้วย ไม่งั้นจะส่งไม่ได้ token นี้เอามาจาก ไปรณ๊ deverlop
	req.Header.Add("Authorization", "Token "+postToken)
	req.Header.Add("Content-Type", "application/json") // postman generate มาไห้ เป็น การ ยิง request api

	res, err := client.Do(req) // postman generate มาไห้ เป็น การ ยิง request api
	if err != nil {            // postman generate มาไห้ เป็น การ ยิง request api
		fmt.Println(err) // postman generate มาไห้ เป็น การ ยิง request api
		return track     // postman generate มาไห้ เป็น การ ยิง request api
	}
	defer res.Body.Close() // postman generate มาไห้ เป็น การ ยิง request api

	body, err := ioutil.ReadAll(res.Body) // postman generate มาไห้ เป็น การ ยิง request api
	if err != nil {                       // postman generate มาไห้ เป็น การ ยิง request api
		fmt.Println(err) // postman generate มาไห้ เป็น การ ยิง request api
		return track     // postman generate มาไห้ เป็น การ ยิง request api
	}

	err = json.Unmarshal(body, &track) // map ข้อมูลที่ได้จาก api ใส่ตัวแปร ที่ มี struct ตามที่เรากำหนดไว้
	if err != nil {
		fmt.Println(err)
		return track
	}
	return track
}
