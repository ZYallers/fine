package fmsg

import (
	"github.com/ZYallers/fine/internal/message"
	"github.com/gin-gonic/gin"
)

type ISender interface {
	Simple(msg interface{}, atAll bool)
	Context(ctx *gin.Context, msg interface{}, reqStr string, stack string, atAll bool)
}

var currentSender ISender

func Sender() ISender {
	if currentSender == nil {
		RegisterSender(new(message.DingTalk))
	}
	return currentSender
}

func RegisterSender(s ISender) {
	currentSender = s
}
