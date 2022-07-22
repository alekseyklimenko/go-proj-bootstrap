package logger

import (
	"bytes"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func Get() *logrus.Logger {
	return log
}

func NewEntry() *logrus.Entry {
	return logrus.NewEntry(log)
}

func WithGinContext(c *gin.Context) *logrus.Entry {
	var l = logrus.NewEntry(log)
	if c.Request.Body != nil {
		b, _ := ioutil.ReadAll(c.Request.Body)
		l = l.WithField("request_body", string(b))
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(b))
	} else {
		c.Set("request_body", nil)
	}
	return l
}
