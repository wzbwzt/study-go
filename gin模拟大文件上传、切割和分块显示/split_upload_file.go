package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
)

func main() {
	r := gin.New()
	r.Use(func(c *gin.Context) {
		defer func() {
			if e := recover(); e != nil {
				log.Error(e.(error))
				c.AbortWithStatusJSON(400, gin.H{"err": e})
			}
		}()
		c.Next()
	})
	r.GET("/get", func(c *gin.Context) {
		c.Request.Header.Set("Transfer-Encoding", "chunked")
		c.Request.Header.Set("Content-type", "image/png")
		for i := 0; i < 5; i++ {
			file, _ := os.Open("./file/img_" + strconv.Itoa(i) + ".jpg")
			data, _ := ioutil.ReadAll(file)
			c.Writer.Write(data)
			c.Writer.Flush()
			time.Sleep(time.Second)
		}
	})
	r.POST("/upload", func(c *gin.Context) {
		fileobj, _ := c.FormFile("file")
		file, _ := fileobj.Open()
		defer file.Close()
		block := fileobj.Size / 5
		index := 0
		for {
			buf := make([]byte, block)
			n, err := file.Read(buf)
			if n == 0 {
				break
			}
			if err != nil && err != io.EOF {
				panic(err)
			}
			saveFile(fmt.Sprintf("img_%d.jpg", index), buf)
			index++
		}
		c.JSON(http.StatusOK, gin.H{"msg": "success"})
	})
	r.Run(":8080")
}

func saveFile(name string, data []byte) {
	save, _ := os.OpenFile("./file/"+name, os.O_CREATE|os.O_RDWR, 6666)
	defer save.Close()
	save.Write(data)
}
