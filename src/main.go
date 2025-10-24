package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	//快速开始
	//quickstart()
	AsciiJSON()
}

func AsciiJSON() {
	router := gin.Default()

	router.GET("/someJson", func(c *gin.Context) {

		// 创建了一个map， key为字符串 ，value是空接口
		// 在 Go 语言中，因为它没有定义任何方法，所以所有类型都默认实现了它。
		// 这意味着 interface{}类型的变量可以容纳任何类型的值，如整数、字符串、切片、甚至是另一个 Map 等
		//。因此，用 map[string]interface{}可以创建一个值类型灵活多变的 Map
		data := map[string]interface{}{
			"lang": "Go语言",
			"tag":  "</br>",
		}
		// AsciiJSON的主要用途是确保生成的JSON数据只包含ASCII字符，
		//它会将非ASCII字符（如中文、表情符号等）转换为其对应的Unicode转义序列（例如 \u5f20\u4e09）。
		//这在某些对字符编码有严格限制的场景下非常有用
		c.AsciiJSON(http.StatusOK, data)
	})

	router.Run(":8080")
}

func quickstart() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.Run(":8080")
	if err != nil {
		return
	}
}
