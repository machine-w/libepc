// Package libepc provides library rfid epc encode model
package libepc

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/hex"
	"encoding/binary"
)

func Encode96bit(source string) (string, int, error) {
	//source string = "j2025022031"
	var target string
	str_len := len(source)
	if str_len < 1 || str_len > 14 {
		return "", -1, errors.New("lenght is error")
	} else if str_len < 8 {
		target = fmt.Sprintf("%d%-7s", str_len, source)
        byteArray := []byte(target)
		target = fmt.Sprintf("%x%x%x%x%x%x%x%x", byteArray[0], byteArray[1],byteArray[2],byteArray[3],byteArray[4],byteArray[5],byteArray[6],byteArray[7])
		return target, str_len, nil
	} else {
		ii, e := strconv.Atoi(source[4 : len(source)-1])
		//fmt.Printf("%x",ii)
		ii = ii << 2
		ii |= ((int(source[3]) - 0x30) & 0x0c) >> 2
		//fmt.Printf("%x",((int(source[3]) - 0x30) & 0x0c) >> 2)
		epc2 := fmt.Sprintf("%08x", ii)
		////////////////////////////////////////////////
		ii = ((int(source[3]) - 0x30) & 0x03)
		ii = (ii << 6) | (int(source[0]) - 0x30)
		ii = (ii << 6) | (int(source[1]) - 0x30)
		ii = (ii << 6) | (int(source[2]) - 0x30)
		ii = (ii << 8) | int(source[str_len-1])
		//fmt.Printf("%x",ii)
		ii = (ii << 4) | str_len
		
		epc1 := fmt.Sprintf("%08x", ii)
		//fmt.Println("")
		target = fmt.Sprintf("%s%s%s%s%s%s%s%s", epc1[6:8], epc1[4:6], epc1[2:4], epc1[0:2], epc2[6:8], epc2[4:6], epc2[2:4], epc2[0:2])
		return target, str_len, e
	}
}
func Decode96bit(source string) (string, int, error) {
	//source string = "j2025022031"
	var target string
	data, err := hex.DecodeString(source)
	if err != nil || len(data) != 8 {
		return "", -1, err
	}
	//fmt.Println(data)
	str_len := int(data[0] & 0x0F )
	//fmt.Println(str_len)
	if str_len < 1 || str_len > 14 {
		return "", -1, errors.New("lenght is error")
	} else if str_len < 8 {
		// for index := 1; index < str_len; index++ {
		// 	target  += fmt.Sprintf("%s",data[index] )
		// }
		target = string(data[1:str_len])
		return target, str_len, nil
	} else {
		low_data := data[4:]
		low_int := binary.LittleEndian.Uint32(low_data)
		//fmt.Printf("%x",low_int)
		four_high := (low_int & 0x03) << 2
		low_int = low_int >> 2
		//fmt.Println(low_int,four_high)
		high_data := data[:4]
		//fmt.Println(high_data)
		high_int := binary.LittleEndian.Uint32(high_data)
		
		high_int = high_int >> 4
		//fmt.Printf("%x",high_int)
		last_code := high_int & 0x0f
		high_int = high_int >> 8
		three_code := high_int & 0x3f
		high_int = high_int >> 6
		two_code := high_int & 0x3f
		high_int = high_int >> 6
		one_code := high_int & 0x3f
		high_int = high_int >> 6
		four_low := high_int & 0x03
		four_code := four_high + four_low
		//fmt.Println("")
		target = fmt.Sprintf("%s%s%s%s%d%d",string(one_code+0x30),string(two_code+0x30),string(three_code+0x30),string(four_code+0x30),low_int,last_code)
		//fmt.Println(last_code,three_code,two_code,one_code)
		return target, str_len, nil
	}
    
}

