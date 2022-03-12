package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
)

//自己实现一个type gin.ResponseWriter interface
type responseWriter struct {
	gin.ResponseWriter
	b *bytes.Buffer
}

//重写Write([]byte) (int, error)
func(w responseWriter) Write(b[]byte) (int, error) {
	//向一个bytes.buffer中再写一份数据
	w.b.Write(b)
	//完成gin.Context.Writer.Write()原有功能
	return w.ResponseWriter.Write(b)
}

func PrintResponse(c *gin.Context) {

	//自己实现一个type gin.ResponseWriter
	writer := responseWriter {
		c.Writer,
		bytes.NewBuffer([]byte{}),
	}
	c.Writer = writer

	c.Next()

	fmt.Println(writer.b.String())
}
