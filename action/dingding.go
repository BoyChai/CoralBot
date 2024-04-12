package action

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	url2 "net/url"

	"github.com/BoyChai/CoralBot/log"
	"github.com/tidwall/gjson"
)

type DingDingHandle struct {
	id          int
	AppKey      string
	AppSecret   string
	accessToken string
}

//var AllHandle []*DingDingHandle

// NewDingDingHandle 创建DingDing动作执行器
func NewDingDingHandle(appKey string, appSecret string) (*DingDingHandle, error) {
	var h DingDingHandle
	h.AppKey = appKey
	h.AppSecret = appSecret
	err := h.getAccessToken()
	if err != nil {
		return nil, err
	}
	return &h, err
}

// GetID 获取当前命令执行器的id
func (h *DingDingHandle) GetID() int {
	return h.id
}

// GetAccessToken 获取token
func (h *DingDingHandle) getAccessToken() (err error) {
	body := make(map[string]interface{})
	body["appKey"] = h.AppKey
	body["appSecret"] = h.AppSecret
	bytesData, err := json.Marshal(body)
	if err != nil {
		return err
	}
	reader := bytes.NewReader(bytesData)
	url := "https://api.dingtalk.com/v1.0/oauth2/accessToken"
	request, err := http.Post(url, "application/json", reader)
	if err != nil {
		return err
	}
	data, err := io.ReadAll(request.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(request.Body)
	if err != nil {
		return err
	}
	if gjson.Get(string(data), "expireIn").String() != "7200" {
		return errors.New(string(data))
	}
	h.accessToken = gjson.Get(string(data), "accessToken").String()
	return nil
}

// Post请求函数
func (h *DingDingHandle) reqPost(url string, msg *bytes.Reader) (statusCode int, bodyData string, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("POST", url, msg)
	if err != nil {
		return 0, "", err
	}

	req.Header.Set("x-acs-dingtalk-access-token", h.accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode, string(body), nil
}

// Get请求函数
func (h *DingDingHandle) reqGet(url string, msg url2.Values) (statusCode int, bodyData string, err error) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, "", err
	}

	req.URL.RawQuery = msg.Encode()

	req.Header.Set("x-acs-dingtalk-access-token", h.accessToken)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, "", err
	}

	return resp.StatusCode, string(body), nil
}

// text 文本类型
type text struct {
	MsgType string `json:"msgType"`
	Content string `json:"content"`
}

func NewTextMsg(content string) string {
	msgInit := text{
		MsgType: "sampleText",
		Content: content,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

// Markdown类型
type markdown struct {
	MsgType string `json:"msgType"`
	Title   string `json:"title"`
	Text    string `json:"text"`
}

func NewMarkdownMsg(title string, text string) string {
	msgInit := markdown{
		MsgType: "sampleMarkdown",
		Title:   title,
		Text:    text,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

// 图片类型
type image struct {
	MsgType  string `json:"msgType"`
	PhotoURL string `json:"photoURL"`
}

func NewImageMsg(photoURL string) string {
	msgInit := image{
		MsgType:  "sampleImageMsg",
		PhotoURL: photoURL,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

// 链接类型
type link struct {
	MsgType    string `json:"msgType"`
	Text       string `json:"text"`
	Title      string `json:"title"`
	PicUrl     string `json:"picUrl"`
	MessageUrl string `json:"messageUrl"`
}

func NewLinkMsg(text string, title string, picUrl string, msgUrl string) string {
	msgInit := link{
		MsgType:    "sampleLink",
		Text:       text,
		Title:      title,
		PicUrl:     picUrl,
		MessageUrl: msgUrl,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

//actionCard类型

type actionCard struct {
	MsgType     string `json:"msgType"`
	Title       string `json:"title"`
	Text        string `json:"text"`
	SingleTitle string `json:"singleTitle"`
	SingleURL   string `json:"singleURL"`
}

func NewActionCardMsg(text string, title string, singleTitle string, singleURL string) string {
	msgInit := actionCard{
		MsgType:     "sampleActionCard",
		Text:        text,
		Title:       title,
		SingleURL:   singleURL,
		SingleTitle: singleTitle,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

type actionCard2 struct {
	MsgType      string `json:"msgType"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	ActionTitle1 string `json:"actionTitle1"`
	ActionURL1   string `json:"actionURL1"`
	ActionTitle2 string `json:"actionTitle2"`
	ActionURL2   string `json:"actionURL2"`
}

func NewActionCard2Msg(text string, title string, actionTitle1 string, actionURL1 string, actionTitle2 string, actionURL2 string) string {
	msgInit := actionCard2{
		MsgType:      "sampleActionCard2",
		Text:         text,
		Title:        title,
		ActionTitle1: actionTitle1,
		ActionURL1:   actionURL1,
		ActionTitle2: actionTitle2,
		ActionURL2:   actionURL2,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

type actionCard3 struct {
	MsgType      string `json:"msgType"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	ActionTitle1 string `json:"actionTitle1"`
	ActionURL1   string `json:"actionURL1"`
	ActionTitle2 string `json:"actionTitle2"`
	ActionURL2   string `json:"actionURL2"`
	ActionTitle3 string `json:"actionTitle3"`
	ActionURL3   string `json:"actionURL3"`
}

func NewActionCard3Msg(text string, title string, actionTitle1 string, actionURL1 string, actionTitle2 string, actionURL2 string, actionTitle3 string, actionURL3 string) string {
	msgInit := actionCard3{
		MsgType:      "sampleActionCard3",
		Text:         text,
		Title:        title,
		ActionTitle1: actionTitle1,
		ActionURL1:   actionURL1,
		ActionTitle2: actionTitle2,
		ActionURL2:   actionURL2,
		ActionTitle3: actionTitle3,
		ActionURL3:   actionURL3,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

type actionCard4 struct {
	MsgType      string `json:"msgType"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	ActionTitle1 string `json:"actionTitle1"`
	ActionURL1   string `json:"actionURL1"`
	ActionTitle2 string `json:"actionTitle2"`
	ActionURL2   string `json:"actionURL2"`
	ActionTitle3 string `json:"actionTitle3"`
	ActionURL3   string `json:"actionURL3"`
	ActionTitle4 string `json:"actionTitle4"`
	ActionURL4   string `json:"actionURL4"`
}

func NewActionCard4Msg(text string, title string, actionTitle1 string, actionURL1 string, actionTitle2 string, actionURL2 string, actionTitle3 string, actionURL3 string, actionTitle4 string, actionURL4 string) string {
	msgInit := actionCard4{
		MsgType:      "sampleActionCard4",
		Text:         text,
		Title:        title,
		ActionTitle1: actionTitle1,
		ActionURL1:   actionURL1,
		ActionTitle2: actionTitle2,
		ActionURL2:   actionURL2,
		ActionTitle3: actionTitle3,
		ActionURL3:   actionURL3,
		ActionTitle4: actionTitle4,
		ActionURL4:   actionURL4,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

type actionCard5 struct {
	MsgType      string `json:"msgType"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	ActionTitle1 string `json:"actionTitle1"`
	ActionURL1   string `json:"actionURL1"`
	ActionTitle2 string `json:"actionTitle2"`
	ActionURL2   string `json:"actionURL2"`
	ActionTitle3 string `json:"actionTitle3"`
	ActionURL3   string `json:"actionURL3"`
	ActionTitle4 string `json:"actionTitle4"`
	ActionURL4   string `json:"actionURL4"`
	ActionTitle5 string `json:"actionTitle5"`
	ActionURL5   string `json:"actionURL5"`
}

func NewActionCard5Msg(text string, title string, actionTitle1 string, actionURL1 string, actionTitle2 string, actionURL2 string, actionTitle3 string, actionURL3 string, actionTitle4 string, actionURL4 string, actionTitle5 string, actionURL5 string) string {
	msgInit := actionCard5{
		MsgType:      "sampleActionCard5",
		Text:         text,
		Title:        title,
		ActionTitle1: actionTitle1,
		ActionURL1:   actionURL1,
		ActionTitle2: actionTitle2,
		ActionURL2:   actionURL2,
		ActionTitle3: actionTitle3,
		ActionURL3:   actionURL3,
		ActionTitle4: actionTitle4,
		ActionURL4:   actionURL4,
		ActionTitle5: actionTitle5,
		ActionURL5:   actionURL5,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

type actionCard6 struct {
	MsgType      string `json:"msgType"`
	Title        string `json:"title"`
	Text         string `json:"text"`
	ButtonTitle1 string `json:"buttonTitle1"`
	ButtonUrl1   string `json:"buttonUrl1"`
	ButtonTitle2 string `json:"buttonTitle2"`
	ButtonUrl2   string `json:"buttonUrl2"`
}

func NewActionCard6Msg(text string, title string, buttonTitle1 string, buttonUrl1 string, buttonTitle2 string, buttonUrl2 string) string {
	msgInit := actionCard6{
		MsgType:      "sampleActionCard6",
		Text:         text,
		Title:        title,
		ButtonTitle1: buttonTitle1,
		ButtonUrl1:   buttonUrl1,
		ButtonTitle2: buttonTitle2,
		ButtonUrl2:   buttonUrl2,
	}
	msg, _ := json.Marshal(msgInit)
	return string(msg)
}

// SendPrivateMessages 批量发送单聊消息
func (h *DingDingHandle) SendPrivateMessages(robotCode string, msg string, userId []string) (statusCode int, bodyData string, err error) {
	// 数据拼接
	body := make(map[string]interface{})
	body["msgParam"] = msg
	body["msgKey"] = gjson.Get(msg, "msgType").String()
	body["userIds"] = userId
	body["robotCode"] = robotCode
	bytesData, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	reader := bytes.NewReader(bytesData)

	return h.reqPost("https://api.dingtalk.com/v1.0/robot/oToMessages/batchSend", reader)
}

// DeletePrivateMessages 批量撤回单聊消息
func (h *DingDingHandle) DeletePrivateMessages(robotCode string, processQueryKeys []string) (statusCode int, bodyData string, err error) {
	// 数据拼接
	body := make(map[string]interface{})
	body["processQueryKeys"] = processQueryKeys
	body["robotCode"] = robotCode
	bytesData, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	reader := bytes.NewReader(bytesData)

	return h.reqPost("https://api.dingtalk.com/v1.0/robot/otoMessages/batchRecall", reader)
}

// QueryPrivateMessages 批量查询机器人单单聊消息是否已读
func (h *DingDingHandle) QueryPrivateMessages(robotCode string, processQueryKeys string) (statusCode int, bodyData string, err error) {
	// 数据拼接
	reader := make(url2.Values)
	reader.Add("robotCode", robotCode)
	reader.Add("processQueryKey", processQueryKeys)

	return h.reqGet("https://api.dingtalk.com/v1.0/robot/oToMessages/readStatus", reader)
}

// SendGroupMessages 企业机器人向内部群发消息
func (h *DingDingHandle) SendGroupMessages(robotCode string, msg string, conversationId string) (statusCode int, bodyData string, err error) {
	// 数据拼接
	body := make(map[string]interface{})
	body["msgParam"] = msg
	body["msgKey"] = gjson.Get(msg, "msgType").String()
	body["openConversationId"] = conversationId
	body["robotCode"] = robotCode
	bytesData, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	reader := bytes.NewReader(bytesData)

	return h.reqPost("https://api.dingtalk.com/v1.0/robot/groupMessages/send", reader)
}

// DeleteGroupMessages 批量撤回群聊消息
func (h *DingDingHandle) DeleteGroupMessages(robotCode string, processQueryKeys []string, conversationId string) (statusCode int, bodyData string, err error) {
	// 数据拼接
	body := make(map[string]interface{})
	body["processQueryKeys"] = processQueryKeys
	body["robotCode"] = robotCode
	body["openConversationId"] = conversationId

	bytesData, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	reader := bytes.NewReader(bytesData)

	return h.reqPost("https://api.dingtalk.com/v1.0/robot/groupMessages/recall", reader)
}

// QueryGroupMessages 查询企业机器人群聊消息用户已读状态
func (h *DingDingHandle) QueryGroupMessages(robotCode string, conversationId string, processQueryKey string, maxResults string, nextToken string) (statusCode int, bodyData string, err error) {
	// 数据拼接
	body := make(map[string]interface{})
	body["openConversationId"] = conversationId
	body["processQueryKey"] = processQueryKey
	body["robotCode"] = robotCode
	body["maxResults"] = maxResults
	body["nextToken"] = nextToken

	bytesData, err := json.Marshal(body)
	if err != nil {
		return 0, "", err
	}

	reader := bytes.NewReader(bytesData)

	return h.reqPost("https://api.dingtalk.com/v1.0/robot/groupMessages/query", reader)
}

// GetHandlerType 实现接口
func (h DingDingHandle) GetHandlerType() string {
	return "DingDIng"
}
