package validator

import (
	"net/http"
)

const (
	Accepted           = "accepted"
	ActiveURL          = "active_url"
	After              = "after"
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
	r       *http.Request
	rule    map[string]string
	message map[string]string
	err     ValidatorError
}

//ValidatorError：存储错误
type ValidatorError struct {
	err map[string][]string
}

//获取错误的第一个验证错误信息
func (err *ValidatorError) First() {
	//TODO:获取校验参数的第一个错误信息
}

func (err *ValidatorError) All() {
	//TODO:获取所有的错误信息
}

//执行数据校验
func (v *Validator) Run() {
	//解析请求
	v.r.ParseForm()
	for filed, ruleStr := v.rule var {
		
	}
}

//拆分对应的校验规则
func (v *Validator) rules(filed, rules string) {
	var error ValidatorError

}
