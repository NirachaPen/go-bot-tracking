package main

type Data struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}
type Event struct {
	Message        Message `json:"message"`
	Mode           bool    `json:"mode"`
	ReplyToken     string  `json:"replyToken"`
	Source         Sorce   `json:"source"`
	Timestamp      string  `json:"timestamp"`
	Type           string  `json:"type"`
	WebhookEventId string  `json:"webhookEventId"`
}
type Message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}
type Sorce struct {
	Type   string `json:"type"`
	UserId string `json:"userId"`
}

type TrackData struct {
	Message  string `json:"message"`
	Response struct {
		Items      map[string]Items `json:"items"`
		TrackCount struct {
			CountNumber     int64  `json:"count_number"`
			TrackCountLimit int64  `json:"track_count_limit"`
			TrackDate       string `json:"track_date"`
		} `json:"track_count"`
	} `json:"response"`
	Status bool `json:"status"`
}
type Items []struct {
	Barcode             string      `json:"barcode"`
	DeliveryDatetime    interface{} `json:"delivery_datetime"`
	DeliveryDescription interface{} `json:"delivery_description"`
	DeliveryStatus      interface{} `json:"delivery_status"`
	Location            string      `json:"location"`
	Postcode            string      `json:"postcode"`
	ReceiverName        interface{} `json:"receiver_name"`
	Signature           interface{} `json:"signature"`
	Status              string      `json:"status"`
	StatusDate          string      `json:"status_date"`
	StatusDescription   string      `json:"status_description"`
}

//ตัวอย่าง response ที่ได้จาก api thai post
// {
//     "response": {
//         "items": {
//             "OB101548343TH": [
//                 {
//                     "barcode": "OB101548343TH",
//                     "status": "103",
//                     "status_description": "รับฝาก",
//                     "status_date": "11/01/2566 10:22:16+07:00",
//                     "location": "วงเวียนใหญ่",
//                     "postcode": "10602",
//                     "delivery_status": null,
//                     "delivery_description": null,
//                     "delivery_datetime": null,
//                     "receiver_name": null,
//                     "signature": null
//                 },
//                 {
//                     "barcode": "OB101548343TH",
//                     "status": "201",
//                     "status_description": "อยู่ระหว่างการขนส่ง",
//                     "status_date": "12/01/2566 03:11:10+07:00",
//                     "location": "ศป.หลักสี่",
//                     "postcode": "10010",
//                     "delivery_status": null,
//                     "delivery_description": null,
//                     "delivery_datetime": null,
//                     "receiver_name": null,
//                     "signature": null
//                 }
//             ]
//         },
//         "track_count": {
//             "track_date": "12/01/2566",
//             "count_number": 3,
//             "track_count_limit": 1000
//         }
//     },
//     "message": "successful",
//     "status": true
// }
