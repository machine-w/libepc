// Package libepc provides library rfid epc encode model
package libepc

import (
	"errors"
	"fmt"
	"strconv"
)

func Encode96bit(source string) (string, int, error) {
	//source string = "j2025022031"
	var target string
	str_len := len(source)
	if str_len < 1 || str_len > 14 {
		return "", -1, errors.New("lenght is error")
	} else if str_len < 8 {
		target = fmt.Sprintf("%d%s", str_len, source)
		return target, str_len, nil
	} else {
		ii, e := strconv.Atoi(source[4 : len(source)-1])
		ii = ii << 2
		ii |= ((int(source[3]) - 0x30) & 0x0c) >> 2
		epc2 := fmt.Sprintf("%08x", ii)
		////////////////////////////////////////////////
		ii = ((int(source[3]) - 0x30) & 0x03)
		ii = (ii << 6) | (int(source[0]) - 0x30)
		ii = (ii << 6) | (int(source[1]) - 0x30)
		ii = (ii << 6) | (int(source[2]) - 0x30)
		ii = (ii << 8) | int(source[str_len-1])
		ii = (ii << 4) | str_len
		epc1 := fmt.Sprintf("%08x", ii)
		target = fmt.Sprintf("%s%s%s%s%s%s%s%s", epc1[6:8], epc1[4:6], epc1[2:4], epc1[0:2], epc2[6:8], epc2[4:6], epc2[2:4], epc2[0:2])
		return target, str_len, e
	}
}
