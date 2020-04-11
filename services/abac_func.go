package services

import "fmt"

type SubAttr struct {
	Name string `json:"name"`
	Role string `json:"role"`
}

type ObjAttr struct {
	Name string `json:"name"`
	Owner string `json:"owner"`
}

type ActAttr struct {
	Name string `json:"name"`
}

type EnvAttr struct {
	Time string `json:"time"`
}

// abac Sub KeyMatch Func
func SubKeyMatch(key1 SubAttr, key2 string) bool {
	fmt.Println("key1:", key1, "key2:", key2)
	return true
}
func SubKeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(SubAttr)
	name2 := args[1].(string)

	return (bool)(SubKeyMatch(name1, name2)), nil
}

// abac Obj KeyMatch Func
func ObjKeyMatch(key1 ObjAttr, key2 string) bool {
	//fmt.Println("key1:", key1, "key2:", key2)
	if key1.Name == key2 {
		return true
	}
	return false
}
func ObjKeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(ObjAttr)
	name2 := args[1].(string)

	return (bool)(ObjKeyMatch(name1, name2)), nil
}

// abac Act KeyMatch Func
func ActKeyMatch(key1 ActAttr, key2 string) bool {
	fmt.Println("key1:", key1, "key2:", key2)
	if key1.Name == key2 {
		return true
	}
	return false
}
func ActKeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(ActAttr)
	name2 := args[1].(string)

	return (bool)(ActKeyMatch(name1, name2)), nil
}

// abac Env KeyMatch Func
func EnvKeyMatch(key1 EnvAttr) bool {
	fmt.Println("key1:", key1)
	return true
}
func EnvKeyMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(EnvAttr)

	return (bool)(EnvKeyMatch(name1)), nil
}