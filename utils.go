package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/axgle/mahonia"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"net/http"
	"strings"
)

func gbkToUtf8(src string) string {
	reader := simplifiedchinese.GB18030.NewDecoder().Reader(strings.NewReader(src))
	bs, err := ioutil.ReadAll(reader)
	if err == nil {
		return string(bs)
	}
	return ""
}

func utf8ToGbk(src string) string {
	enc := mahonia.NewEncoder("gbk")
	return enc.ConvertString(src)
}

func httpPost(serverAddr string, data interface{}) {
	client := http.DefaultClient
	if bs, err := json.Marshal(data); err == nil {
		client.Post(serverAddr, "application//json", bytes.NewReader(bs))
	}
}

func getBase64(data []byte) string {
	strBytes := []byte(data)
	encoded := base64.StdEncoding.EncodeToString(strBytes)
	return encoded
}
