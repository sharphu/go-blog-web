package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Setting struct {
	vp *viper.Viper
}

func NewSetting() (*Setting, error){
	vp := viper.New()
	vp.SetConfigName("config")
	vp.AddConfigPath("configs/")
	vp.SetConfigType("yaml")
	err := vp.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{vp}
	s.WatchSettingChange()

	return s, nil
}

//func NewSetting(configs ...string) (*Setting, error){
//	/*
//	从命令行读取配置文件路径
//	 */
//	vp := viper.New()
//	vp.SetConfigName("config")
//	for _, config := range configs {
//		if config != "" {
//			vp.AddConfigPath(config)
//		}
//	}
//	vp.SetConfigType("yaml")
//	err := vp.ReadInConfig()
//	if err != nil {
//		return nil, err
//	}
//	s := &Setting{vp}
//	s.WatchSettingChange()
//
//	return s, nil
//}

func (s *Setting) WatchSettingChange() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(in fsnotify.Event) {
			_ = s.ReloadAllSection()
		})
	}()
}
