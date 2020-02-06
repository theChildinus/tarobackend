package utils

import (
	"os"
	"reflect"
)

type Page struct {
	PageIndex int64 `json:"page_index"`
	PageSize  int64 `json:"page_size"`
}

type NormalResp struct {
	Code   int    `json:"code"`
	NorMsg string `json:"normsg"`
}

type ErrorResp struct {
	Code   int    `json:"code"`
	ErrMsg string `json:"errmsg"`
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

func FileExistAndCreate(path string) (bool, error) {
	if _, err := os.Stat(path); err == nil {
		return true, nil
	} else if os.IsNotExist(err) {
		if _, err := os.Create(path); err != nil {
			return false, err
		} else {
			return true, nil
		}
	} else {
		return false, nil
	}
}
