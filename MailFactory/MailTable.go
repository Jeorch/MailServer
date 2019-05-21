package MailFactory

import (
	"Mail/MailHandler"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmMongodb"
	"github.com/alfredyang1986/BmServiceDef/BmDaemons/BmRedis"
)

type MailTable struct{}

var MAIL_MODEL_FACTORY = map[string]interface{}{
}

var MAIL_STORAGE_FACTORY = map[string]interface{}{
}

var MAIL_RESOURCE_FACTORY = map[string]interface{}{
}

var MAIL_FUNCTION_FACTORY = map[string]interface{}{
	"MailCommonPanicHandle":	MailHandler.CommonPanicHandle{},
	"MailSendHandler":		MailHandler.MailSendHandler{},
}
var MAIL_MIDDLEWARE_FACTORY = map[string]interface{}{
}

var MAIL_DAEMON_FACTORY = map[string]interface{}{
	"BmMongodbDaemon": BmMongodb.BmMongodb{},
	"BmRedisDaemon":   BmRedis.BmRedis{},
}

func (t MailTable) GetModelByName(name string) interface{} {
	return MAIL_MODEL_FACTORY[name]
}

func (t MailTable) GetResourceByName(name string) interface{} {
	return MAIL_RESOURCE_FACTORY[name]
}

func (t MailTable) GetStorageByName(name string) interface{} {
	return MAIL_STORAGE_FACTORY[name]
}

func (t MailTable) GetDaemonByName(name string) interface{} {
	return MAIL_DAEMON_FACTORY[name]
}

func (t MailTable) GetFunctionByName(name string) interface{} {
	return MAIL_FUNCTION_FACTORY[name]
}

func (t MailTable) GetMiddlewareByName(name string) interface{} {
	return MAIL_MIDDLEWARE_FACTORY[name]
}
