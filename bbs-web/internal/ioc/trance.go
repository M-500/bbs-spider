package ioc

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
	"strings"
)

// @Description
// @Author 代码小学生王木木
// @Date 2024-03-28 12:06

var InstanceTrans ut.Translator

func InitTrans(cfg *Config) ut.Translator {

	//修改gin框架中的validator引擎属性, 实现定制
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//注册一个获取json的tag的自定义方法
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})

		zhT := zh.New() //中文翻译器
		enT := en.New() //英文翻译器
		//第一个参数是备用的语言环境，后面的参数是应该支持的语言环境
		uni := ut.New(enT, zhT, zhT)
		InstanceTrans, ok = uni.GetTranslator(cfg.Language)
		if !ok {
			//return fmt.Errorf("uni.GetTranslator(%s)", locale)
			panic(any(fmt.Errorf("uni.GetTranslator(%s)", cfg.Language)))
		}

		switch cfg.Language {
		case "en":
			entranslations.RegisterDefaultTranslations(v, InstanceTrans)
		case "zh":
			zhtranslations.RegisterDefaultTranslations(v, InstanceTrans)
		default:
			entranslations.RegisterDefaultTranslations(v, InstanceTrans)
		}
		return InstanceTrans
	}
	return nil
}
