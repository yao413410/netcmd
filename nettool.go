package netcmd

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

//BytesCombine 多个[]byte数组合并成一个[]byte
func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

//BytesToInt isSymbol表示有无符号
func BytesToInt(b []byte, isSymbol bool) (int, error) {
	if isSymbol {
		return bytesToIntS(b)
	}
	return bytesToIntU(b)
}

//BytesToInt64 isSymbol表示有无符号
func BytesToInt64(b []byte, isSymbol bool) (int64, error) {
	if isSymbol {
		return bytesToIntS64(b)
	}
	return bytesToIntU64(b)
}

//字节数(大端)组转成int(无符号的)
func bytesToIntU(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp uint8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp uint16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp uint32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//bytesToIntU64 字节数(大端)组转成int64(无符号的)
func bytesToIntU64(b []byte) (int64, error) {
	if len(b) == 8 {
		bytesBuffer := bytes.NewBuffer(b)
		var tmp uint64
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int64(tmp), err
	}
	return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
}

//bytesToIntS 字节数(大端)组转成int(有符号)
func bytesToIntS(b []byte) (int, error) {
	if len(b) == 3 {
		b = append([]byte{0}, b...)
	}
	bytesBuffer := bytes.NewBuffer(b)
	switch len(b) {
	case 1:
		var tmp int8
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 2:
		var tmp int16
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	case 4:
		var tmp int32
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return int(tmp), err
	default:
		return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
	}
}

//bytesToIntS64 字节数(大端)组转成int(有符号)
func bytesToIntS64(b []byte) (int64, error) {
	if len(b) == 8 {
		bytesBuffer := bytes.NewBuffer(b)
		var tmp int64
		err := binary.Read(bytesBuffer, binary.BigEndian, &tmp)
		return tmp, err
	}
	return 0, fmt.Errorf("%s", "BytesToInt bytes lenth is invaild!")
}

//IntToBytes 整形转换成字节
func IntToBytes(n int, b byte) ([]byte, error) {
	switch b {
	case 1:
		tmp := int8(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 2:
		tmp := int16(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil
	case 3, 4:
		tmp := int32(n)
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.BigEndian, &tmp)
		return bytesBuffer.Bytes(), nil

	}
	return nil, fmt.Errorf("IntToBytes b param is invaild")
}

//IntToBytes64 整形转换成字节
func IntToBytes64(n int64) ([]byte, error) {
	tmp := int64(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, &tmp)
	return bytesBuffer.Bytes(), nil
}

//Printf 打印错误
func PrintfWarning(format string, a ...interface{}) {
	s := fmt.Sprintf(format, a...)
	fmt.Printf("%c[1;40;33m%s%c[0m\n", 0x1B, s, 0x1B)
}

