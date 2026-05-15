package huaweisms

import "encoding/xml"

// Message 表示一条短信
type Message struct {
	Phone   string `xml:"Phone"`
	Content string `xml:"Content"`
	Date    string `xml:"Date"`
}

// 内部使用的结构（不导出）
type stateResponse struct {
	State        string `xml:"State"`
	PasswordType int    `xml:"password_type"`
}

type loginRequest struct {
	XMLName      xml.Name `xml:"request"`
	Username     string   `xml:"Username"`
	Password     string   `xml:"Password"`
	PasswordType int      `xml:"password_type"`
}

type smsRequest struct {
	XMLName         xml.Name `xml:"request"`
	PageIndex       int      `xml:"PageIndex"`
	ReadCount       int      `xml:"ReadCount"`
	BoxType         int      `xml:"BoxType"`
	SortType        int      `xml:"SortType"`
	Ascending       int      `xml:"Ascending"`
	UnreadPreferred int      `xml:"UnreadPreferred"`
}

type smsResponse struct {
	Messages struct {
		Message []Message `xml:"Message"`
	} `xml:"Messages"`
}
