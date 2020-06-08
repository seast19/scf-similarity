package main

import "testing"

// func TestMy(t *testing.T) {
// 	MyTest("我是谁", "我是你")
// }

func TestPost(t *testing.T) {
	d := DefineEvent{}
	d.Body = `{"text":[{"id":1,"first":"我爱中国","second":"我爱祖国"}]}`
	Scf(d)
}
