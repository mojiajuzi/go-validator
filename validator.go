package validator

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"
	"strings"
)

//定义校验规则
const (
	Accepted           = "accepted"
	ActiveURL          = "active_url"
	After              = "after"
	Alpha              = "alpha"
	AlphaDash          = "alpha_dash"
	AlphaNum           = "alpha_num"
	Array              = "array"
	Befor              = "befor"
	Between            = "between"
	Boolean            = "boolean"
	Confirmed          = "confirmed"
	Date               = "date"
	DateFormat         = "date_format"
	Different          = "different"
	Digits             = "digits"
	DigitsBetween      = "digits_between"
	Email              = "email"
	Exists             = "exists"
	Image              = "image"
	In                 = "in"
	Integer            = "integer"
	IP                 = "ip"
	JSON               = "json"
	Max                = "max"
	Min                = "min"
	Mimes              = "mimes"
	NotIn              = "not_in"
	Numeric            = "numeric"
	Regex              = "regex"
	Required           = "required"
	RequiredIf         = "required_if"
	RequiredUnless     = "required_uness"
	RequiredWith       = "required_with"
	RequiredWithAll    = "required_with_all"
	RequiredWithout    = "required_without"
	RequiredWithoutAll = "required_without_all"
	Same               = "same"
	Size               = "size"
	String             = "string"
	TimeZone           = "timezone"
	Unique             = "unique"
	URL                = "url"
)

//Validator数据校验的对象
type Validator struct {
	r                *http.Request
	rule             map[string]string
	message          map[string]string
	err              map[string][]string
	errConfigMessage map[string]map[string]interface{}
}

//First 获取错误的第一个验证错误信息
func (v *Validator) First() {
	//TODO:获取校验参数的第一个错误信息
}

//All 获取所有的错误信息
func (v *Validator) All() {
	//TODO:获取所有的错误信息
}

//Run 执行数据校验
func (v *Validator) Run() {
	//解析请求
	v.r.ParseForm()
	//加载文件替换信息

	for field, ruleStr := range v.rule {
		value := v.r.FormValue(field)
		rules := strings.Split(strings.Trim(ruleStr, ""), "|")
		v.subRuleValidator(value, field, rules)
	}
}

//loadConfigErrorMessage 加载错误配置信息
func (v *Validator) loadConfigErrorMessage(lang string) {
	//TODO:错误信息信息处理
	f, err := ioutil.ReadFile("data.json")
	if err != nil {
		fmt.Println("load config gile error")
	}
	err = json.Unmarshal(f, &v.errConfigMessage)
	if err != nil {
		fmt.Println(err.Error())
	}
}

//拆分对应的校验规则
func (v *Validator) subRuleValidator(value interface{}, filed string, rules []string) {
	for _, rule := range rules {
		switch rule {
		case Accepted:
		case ActiveURL:
		case Alpha:
		}
	}
}

//validatorAccepted 校验是否同意的布尔值
func (v *Validator) validatorAccepted(field string, value interface{}) {
	re := reflect.TypeOf(value).Kind()
	if re != reflect.Bool {
		v.err[field] = append(v.err[field], "the "+field+" type must be bool")
	}
}
