package req

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	chTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/lstack-org/go-web-framework/pkg/code"
	k8sErrs "k8s.io/apimachinery/pkg/util/errors"
	"strings"
)

var (
	ZhTrans ut.Translator
	EnTrans ut.Translator
	Val               = validator.New()
)

func init() {
	zhT := zh.New()
	enT := en.New()
	uni := ut.New(enT, zhT, enT)

	ZhTrans, _ = uni.GetTranslator(zhT.Locale())
	_ = chTranslations.RegisterDefaultTranslations(Val, ZhTrans)

	EnTrans, _ = uni.GetTranslator(enT.Locale())
	_ = enTranslations.RegisterDefaultTranslations(Val, EnTrans)
	return
}

func translate(ctx *gin.Context, errs ...error) string {
	var newErrs []error
	for _, err := range errs {
		validationErr, ok := err.(validator.ValidationErrors)
		if !ok {
			return err.Error()
		}

		lang := ctx.GetHeader(code.AcceptLanguageHeader)
		switch lang {
		case code.AcceptLanguageZh, ZhTrans.Locale():
			newErrs = append(newErrs, errors.New(validationErrorsTranslationsToString(validationErr.Translate(ZhTrans))))
		default:
			newErrs = append(newErrs, errors.New(validationErrorsTranslationsToString(validationErr.Translate(EnTrans))))
		}
	}
	return k8sErrs.NewAggregate(newErrs).Error()
}

func validationErrorsTranslationsToString(err validator.ValidationErrorsTranslations) (errMsg string) {
	var errs []string
	for k, v := range err {
		errs = append(errs, fmt.Sprintf("%s:%s", k, v))
	}
	errMsg = strings.Join(errs, ";")
	return
}
