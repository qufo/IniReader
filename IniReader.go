package IniReader

import (
	"io/ioutil"
	"strings"
)

type IniReader struct {
	file 	string
	items 	map[string]string
}
// 创建，file - ini 文件名称
func NewIniReader(file string) *IniReader {
	buf,err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	items := make(map[string]string)
	confArr := strings.Split(string(buf),"\n")
	for _,lines := range confArr {
		line := strings.Split(lines,"=")
		if len(line) == 2 {
			key := strings.TrimSpace(line[0])
			value := strings.TrimSpace(strings.Split(line[1],"#")[0])
			items[key] = value
		}
	}
	return &IniReader{file,items}
}

// 取得 ini 中的值
func (this *IniReader) Get (key string ) string {
	if _,ok := this.items[key]; ok {
		return this.items[key]
	}
	return ""
}
