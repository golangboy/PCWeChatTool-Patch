package main

/*
	#include "hook.h"
	#include "string.h"
	#include "stdlib.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/**
 * @Description:  当接收到消息时(包括所有文字、语音、群、等其他消息)
 * @param wxId	  发起者微信ID,当消息是消息时是群ID 形如@chatom
 * @param wxMsg	  消息正文
 * @param msgType 消息类型 msgType=1 时 是文字消息,否则是一个XML结构
 */
//export recvMsg
func recvMsg(wxId string, wxMsg string, msgType int) {
	wxMsg = gbkToUtf8(wxMsg)
	wxId = gbkToUtf8(wxId)
	msgT := msgType
	fmt.Println(wxMsg, msgT)
}

/**
 * @Description:  当接收到图片消息时
 * @param wxId    发送者微信ID
 * @param imgPath 图片的本地绝对路径
 */
//export recvImageMsg
func recvImageMsg(wxId string, imgPath string) {
	imgPath = gbkToUtf8(imgPath)
	wxId = gbkToUtf8(wxId)
	fmt.Println(imgPath, wxId)
}

/**
 * @Description:        当接收到语音消息时
 * @param wxId          发送者微信ID
 * @param ptrVoiceData  语音数据指针
 * @param dataLen		语音数据长度
 */
//export recvVoiceMsg
func recvVoiceMsg(wxId string, ptrVoiceData unsafe.Pointer, dataLen int) {
	wxId = gbkToUtf8(wxId)
	voiceLen := dataLen
	voiceBuff := make([]byte, dataLen)
	C.memcpy(unsafe.Pointer(&voiceBuff[0]), ptrVoiceData, C.uint(dataLen))
	fmt.Println(voiceLen)
}

/**
 * @Description: 初始函数 DLL被载入时会被调用
 */
func init() {
	C.StartHook()
	//go func() {
	//	g := gin.Default()
	//	g.Handle("GET", "/text", func(context *gin.Context) {
	//		context.Writer.WriteString(wxid + " " + txt + " " + strconv.Itoa(msgT))
	//	})
	//	g.Handle("GET", "/img", func(context *gin.Context) {
	//		context.Writer.WriteString(wxid + " " + img)
	//	})
	//	g.Handle("GET", "/voice", func(context *gin.Context) {
	//		context.Writer.Write(voiceBuff)
	//	})
	//	g.Run(":8080")
	//}()
}

/**
 * @Description: Go生成动态链接库需要的入口函数
 */
func main() {

}

/**
 * @Description: 用于WeChat.exe添加的导入函
 */
//export Enter
func Enter() {

}
