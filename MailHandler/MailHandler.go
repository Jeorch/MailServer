package MailHandler

import (
	"encoding/json"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
	"github.com/go-gomail/gomail"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"reflect"
	"strconv"
)

type MailSendHandler struct {
	Method     string
	HttpMethod string
	Args       []string
	db         *BmMongodb.BmMongodb
	rd         *BmRedis.BmRedis
}

type email struct {
	Email string `json:"email"`
	Subject string	`json:"subject"`
	Content string 	`json:"content"`
	ContentType string	`json:"content-type"`
}

func (h MailSendHandler) NewMailHandler(args ...interface{}) MailSendHandler {
	var m *BmMongodb.BmMongodb
	var r *BmRedis.BmRedis
	var hm string
	var md string
	var ag []string
	for i, arg := range args {
		if i == 0 {
			sts := arg.([]BmDaemons.BmDaemon)
			for _, dm := range sts {
				tp := reflect.ValueOf(dm).Interface()
				tm := reflect.ValueOf(tp).Elem().Type()
				if tm.Name() == "BmMongodb" {
					m = dm.(*BmMongodb.BmMongodb)
				} else if tm.Name() == "BmRedis" {
					r = dm.(*BmRedis.BmRedis)
				}
			}
		} else if i == 1 {
			md = arg.(string)
		} else if i == 2 {
			hm = arg.(string)
		} else if i == 3 {
			lst := arg.([]string)
			for _, str := range lst {
				ag = append(ag, str)
			}
		} else {
		}
	}

	return MailSendHandler{Method: md, HttpMethod: hm, Args: ag, db: m, rd: r}
}

func (h MailSendHandler) SendMail(w http.ResponseWriter, r *http.Request, _ httprouter.Params) int {
	w.Header().Add("Content-Type", "application/json")

	body, err := ioutil.ReadAll(r.Body)
	response := map[string]interface{}{}

	if err != nil {
		response["status"] = "error"
		response["msg"] = "request解析失败"
		enc := json.NewEncoder(w)
		enc.Encode(response)
		return 1
	}
	toEmail := email{}
	err = json.Unmarshal(body, &toEmail)

	if err != nil {
		response["status"] = "error"
		response["msg"] = "Json解析失败"
		enc := json.NewEncoder(w)
		enc.Encode(response)
		return 1
	}

	if len(toEmail.Email) == 0 {
		response["status"] = "error"
		response["msg"] = "Email为空"
		enc := json.NewEncoder(w)
		enc.Encode(response)
		return 1
	}

	m := gomail.NewMessage()
	m.SetHeader("From", h.Args[0])
	m.SetHeader("To", toEmail.Email)
	m.SetHeader("Subject", toEmail.Subject)
	m.SetBody(toEmail.ContentType, toEmail.Content)

	port, _ := strconv.Atoi(h.Args[3])

	d := gomail.NewDialer(h.Args[2], port, h.Args[0], h.Args[1])
	//gomail.NewDialer(h.Args[2], port, h.Args[0], h.Args[1])
	d.DialAndSend(m)
	//if err := d.DialAndSend(m); err != nil {
	//	response["status"] = "error"
	//	response["msg"] = "邮件发送失败"
	//	enc := json.NewEncoder(w)
	//	enc.Encode(response)
	//}

	response["status"] = "success"
	response["msg"] = "邮件发送成功"
	enc := json.NewEncoder(w)
	enc.Encode(response)

	return 0
}

func (h MailSendHandler) GetHttpMethod() string {
	return h.HttpMethod
}

func (h MailSendHandler) GetHandlerMethod() string {
	return h.Method
}
