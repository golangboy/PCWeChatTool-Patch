package main

import (
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
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
