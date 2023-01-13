package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// token สำหรับ ส่งข้อความ line
var token = ""

// token สำหรับ ยิง api thai post
var postToken = ""

func main() {
	r := gin.New()
	r.POST("/hook", HookMessage)
	r.Run(":8000")
}
func HookMessage(c *gin.Context) {
	var d Data
	err := c.ShouldBind(&d)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	if d.Events[0].Message.Type == "text" {
		tracking := Tracking(d.Events[0].Message.Text) // รับค่าที่ได้จาก การ ยิง api ไปยัง ไปรศนี
		SendMessage(d.Events[0].Source.UserId, d.Events[0].Message.Text, tracking.Response.Items[d.Events[0].Message.Text])

	}
	c.JSON(200, gin.H{
		"message": "ok",
	})
}
