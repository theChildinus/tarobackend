package utils

import (
	"reflect"
)

type Page struct {
	PageIndex int64 `json:"page_index"`
	PageSize  int64 `json:"page_size"`
}

type NormalResp struct {
	Code   int
	NorMsg string
}

type ErrorResp struct {
	Code   int
	ErrMsg string
}

func BuildJsonResp(c interface{}, msgType, msg string) {
	e := reflect.ValueOf(c)
	if msgType == "Normal" {
		nr := &NormalResp{Code: 0, NorMsg: msg}
		e.Elem().FieldByName("Data").SetMapIndex(reflect.ValueOf("json"), reflect.ValueOf(nr))
	} else if msgType == "Error" {
		er := &ErrorResp{Code: -1, ErrMsg: msg}
		e.Elem().FieldByName("Data").SetMapIndex(reflect.ValueOf("json"), reflect.ValueOf(er))
	}
	e.MethodByName("ServeJSON").CallSlice([]reflect.Value{reflect.ValueOf([]bool{})})
}
