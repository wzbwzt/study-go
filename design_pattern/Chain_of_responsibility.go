package main

import "fmt"

//职责链

// Request 定义请求结构体
type Request struct {
	username string
	password string
	token    string
}

// Validator 定义职责链接口
type Validator interface {
	SetNext(Validator) Validator
	Validate(Request) error
}

// BaseValidator 基础验证器
type BaseValidator struct {
	nextValidator Validator
}

func (b *BaseValidator) SetNext(v Validator) Validator {
	b.nextValidator = v
	return v
}

func (b *BaseValidator) Validate(r Request) error {
	if b.nextValidator != nil {
		return b.nextValidator.Validate(r)
	}
	return nil
}

// UsernameValidator 用户名验证器
type UsernameValidator struct {
	BaseValidator
}

func (u *UsernameValidator) Validate(r Request) error {
	if r.username == "" {
		return fmt.Errorf("Username cannot be empty")
	}
	return u.BaseValidator.Validate(r)
}

// PasswordValidator 密码验证器
type PasswordValidator struct {
	BaseValidator
}

func (p *PasswordValidator) Validate(r Request) error {
	if r.password == "" {
		return fmt.Errorf("Password cannot be empty")
	}
	return p.BaseValidator.Validate(r)
}

// TokenValidator Token验证器
type TokenValidator struct {
	BaseValidator
}

func (t *TokenValidator) Validate(r Request) error {
	if r.token == "" {
		return fmt.Errorf("Token cannot be empty")
	}
	return t.BaseValidator.Validate(r)
}

func main() {
	// 创建验证器链
	usernameValidator := &UsernameValidator{}
	passwordValidator := &PasswordValidator{}
	tokenValidator := &TokenValidator{}

	// 设置验证器链
	usernameValidator.SetNext(passwordValidator).SetNext(tokenValidator)

	// 创建请求
	request := Request{username: "testuser", password: "testpass", token: "testtoken"}

	// 验证请求
	if err := usernameValidator.Validate(request); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Request is valid")
	}
}
