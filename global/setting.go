package global

import (
	"blog-web/pkg/logger"
	"blog-web/pkg/setting"
)

var (
	ServerSetting *setting.ServerSettingS
	AppSetting *setting.AppSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger *logger.Logger
	JWTSetting *setting.JWTSettingS
	EmailSetting *setting.EmailSettingS
)
