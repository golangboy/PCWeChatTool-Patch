package main

/*
	#include "hook.h"
	#include "string.h"
	#include "stdlib.h"
*/
import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zyx4843/gojson"
	"strconv"
	"unsafe"
)

const defaultPostServer = "http://localhost:19730"
const defaultBindPort = 8888

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
	httpPost(defaultPostServer+"/msg", struct {
		WxId string `json:"wx_id"`
		Msg  string `json:"msg"`
		Type int    `json:"type"`
	}{wxId, wxMsg, msgType})
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
	httpPost(defaultPostServer+"/img", struct {
		WxId            string `json:"wx_id"`
		ImgLocalAbsPath string `json:"img_local_abs_path"`
	}{wxId, imgPath})
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
	voiceBuff := make([]byte, dataLen)
	C.memcpy(unsafe.Pointer(&voiceBuff[0]), ptrVoiceData, C.uint(dataLen))

	httpPost(defaultPostServer+"/voice", struct {
		WxId      string `json:"wx_id"`
		VoiceData string `json:"voice_data"`
	}{wxId, getBase64(voiceBuff)})
}

/**
 * @Description: 启动本地HTTP服务提供接口
 */
func startSendMsgServer() {
	g := gin.Default()
	g.Handle("POST", "text", func(context *gin.Context) {
		if bs, err := context.GetRawData(); err == nil {
			jsonData := string(bs)
			wxId := gojson.Json(jsonData).Get("wx_id").Tostring()
			textMsg := gojson.Json(jsonData).Get("text").Tostring()
			fmt.Println(wxId, textMsg)

			if len(wxId) > 0 && len(textMsg) > 0 {
				CSwxId := C.CString(utf8ToGbk(wxId))
				CSMsg := C.CString(utf8ToGbk(textMsg))
				C.SendTextMessage(CSwxId, CSMsg)
				defer C.free(unsafe.Pointer(CSwxId))
				defer C.free(unsafe.Pointer(CSMsg))
			}
		}
	})
	g.Handle("POST", "file", func(context *gin.Context) {
		if bs, err := context.GetRawData(); err == nil {
			jsonData := string(bs)
			wxId := gojson.Json(jsonData).Get("wx_id").Tostring()
			filePath := gojson.Json(jsonData).Get("path").Tostring()
			if len(wxId) > 0 && len(filePath) > 0 {
				CSwxId := C.CString(utf8ToGbk(wxId))
				CSPath := C.CString(utf8ToGbk(filePath))

				C.SendFileMessage(CSwxId, CSPath)

				defer C.free(unsafe.Pointer(CSwxId))
				defer C.free(unsafe.Pointer(CSPath))
			}
		}
	})
	g.Run(":" + strconv.Itoa(defaultBindPort))
}

/**
 * @Description: 初始函数 DLL被载入时会被调用
 */
func init() {
	C.StartHook()
	go func() {
		startSendMsgServer()
	}()
}

/**
 * @Description: Go生成动态链接库需要的入口函数
 */
func main() {

}

/**
 * @Description: 用于WeChat.exe添加的导入函数
 */
//export Enter
func Enter() {

}
