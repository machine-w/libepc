// Package libepc provides rfid epc encode and decode
package libepc

import (
	"testing"
)

func Test_Encode96bit_err(t *testing.T) {
	if s, l, e := Encode96bit(""); s != "" || l != -1 || e == nil {
		t.Error("长度不合法空字符串测试不通过")
	} else if s, l, e := Encode96bit("td12345678900000000000000"); s != "" || l != -1 || e == nil {
		t.Error("长度不合法超长测试不通过")
	}
}
func Test_Encode96bit_eight(t *testing.T) {
	if s, l, e := Encode96bit("12345"); s != "512345" || l != 5 || e != nil {
		t.Error("小于8个字符测试不通过")
	} else if s, l, e := Encode96bit("1234567"); s != "71234567" || l != 7 || e != nil {
		t.Error("7个字符测试不通过")
	}
}
func Test_Encode96bit_upeight(t *testing.T) {
	if s, l, e := Encode96bit("TD002349754"); s != "4b0350247c570e00" || l != 11 || e != nil {
		t.Error("图书td号测试不通过")
	} else if s, l, e := Encode96bit("j2025022031"); s != "1b0308baeca61e00" || l != 11 || e != nil {
		t.Error("层架标测试不通过")
	} else if _, _, e := Encode96bit("J202502T031"); e == nil {
		t.Error("4位后数字转换测试不通过")
	}
}
