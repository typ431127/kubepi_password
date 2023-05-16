package model

import "time"

type BaseModel struct {
	ApiVersion string    `json:"apiVersion"`
	Kind       string    `json:"kind"`
	CreateAt   time.Time `json:"createAt" storm:"index"`
	UpdateAt   time.Time `json:"updateAt" storm:"index"`
	BuiltIn    bool      `json:"builtIn"`
	CreatedBy  string    `json:"createdBy"`
}

type Metadata struct {
	Name        string `json:"name" storm:"unique" `
	Description string `json:"description"`
	UUID        string `json:"uuid" storm:"id,index,unique"`
}

type Authenticate struct {
	Password string `json:"password"`
	Token    string `json:"token"`
}

type Mfa struct {
	Enable bool   `json:"enable"`
	Secret string `json:"secret"`
}

type User struct {
	BaseModel    `storm:"inline"`
	Metadata     `storm:"inline"`
	NickName     string       `json:"nickName" storm:"index"`
	Email        string       `json:"email" storm:"unique"`
	Language     string       `json:"language"`
	IsAdmin      bool         `json:"isAdmin"`
	Authenticate Authenticate `json:"authenticate"`
	Type         string       `json:"type"`
	Mfa          Mfa          `json:"mfa"`
}
