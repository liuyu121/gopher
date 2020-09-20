package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

// 定义 tag 名称
const tagname = "validate"

// 定义电子邮件的正则表达式
var mailRex = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)

// 定义一个 interface
type Validator interface {
	Validate(interface{}) (bool, error)
}

// 默认的空校验器
type DefaultValidator struct {
}

// 实现了 Validator 接口的 DefaultValidator
func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

// 定义整数的校验器
type NumberValidator struct {
	Min int
	Max int
}

func (v NumberValidator) Validate(val interface{}) (bool, error) {
	// interface 转化成 int 类型
	num := val.(int)

	if v.Min > v.Max {
		return false, fmt.Errorf("illegal min=%v and max=%v value", v.Min, v.Max)
	}
	if num < v.Min {
		return false, fmt.Errorf("must be greater than %v", v.Min)
	}
	if v.Max >= v.Min && num > v.Max {
		return false, fmt.Errorf("must be less than %v", v.Max)
	}

	return true, nil
}

// 定义 string 的 校验器
type StringValidator struct {
	Min int
	Max int
}

func (v StringValidator) Validate(val interface{}) (bool, error) {
	l := len(val.(string))

	if l == 0 {
		return false, fmt.Errorf("cannot be blank")
	}

	if l < v.Min {
		return false, fmt.Errorf("must be at least %v chars long", v.Min)
	}

	if v.Max >= v.Min && l > v.Max {
		return false, fmt.Errorf("should be less than %v chars long", v.Max)
	}

	return true, nil
}

// 定义 email 的校验器
type EmailValidator struct {
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	if !mailRex.MatchString(val.(string)) {
		return false, fmt.Errorf("is not a valid email address")
	}

	return true, nil
}

// 获取 tag 对应的 校验器
func getValidatorFromTag(tag string) Validator {
	args := strings.Split(tag, ",")

	switch args[0] {
	case "number":
		validator := NumberValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "string":
		validator := StringValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d", &validator.Min, &validator.Max)
		return validator
	case "email":
		return EmailValidator{}
	}

	return DefaultValidator{}
}

func validateStruct(s interface{}) []error {
	var errs []error

	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get(tagname)

		// 跳过
		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)
		// 这里获取 interface{} 类型的值，因为 validator 接收的是 interface{} 类型的值
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))
		}
	}

	return errs
}

// 如果需要 reflect 访问，这里字段必须是 public 的，否则
// panic: reflect.Value.Interface: cannot return value obtained from unexported field or method
type Admin2 struct {
	Id    int    `validate:"number,min=1,max=1000"`
	Name  string `validate:"string,min=2,max=10"`
	Bio   string `validate:"string"`
	Email string `validate:"email"`
}

func main() {
	user := Admin2{
		88, "superlongstring", "", "1@qq.com",
	}

	fmt.Println("errors： ")

	for i, err := range validateStruct(user) {
		fmt.Printf("\t%d. %s\n", i+1, err.Error())
	}
}
