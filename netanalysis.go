package netcmd

import "fmt"

//CmdData cmd数据
type CmdData struct {
	data  []byte
	index int
}

//InitData 初始化数据
func (p *CmdData) InitData(data []byte) {
	if data == nil {
		p.data = p.data[0:0]
	} else {
		p.data = data
	}
	p.index = 0
}

//SetIndex 设置下标
func (p *CmdData) SetIndex(index int) {
	p.index = index
}

//GetInt 获取int
func (p *CmdData) GetInt() int {
	num := len(p.data)
	if num < p.index+4 {
		return 0
	}
	n, err := BytesToInt(p.data[p.index:p.index+4], false)
	if err != nil {
		return 0
	}
	p.index += 4
	return n
}

//GetInt8 获取int8
func (p *CmdData) GetInt8() int {
	num := len(p.data)
	if num < p.index+1 {
		return 0
	}
	n, err := BytesToInt(p.data[p.index:p.index+1], false)
	if err != nil {
		return 0
	}
	p.index++
	return n
}

//GetInt16 获取int16
func (p *CmdData) GetInt16() int {
	num := len(p.data)
	if num < p.index+2 {
		return 0
	}
	n, err := BytesToInt(p.data[p.index:p.index+2], false)
	if err != nil {
		return 0
	}
	p.index += 2
	return n
}

//GetInt64 获取int64
func (p *CmdData) GetInt64() int64 {
	num := len(p.data)
	if num < p.index+8 {
		return 0
	}
	n, err := BytesToInt64(p.data[p.index:p.index+8], false)
	if err != nil {
		return 0
	}
	p.index += 8
	return n
}

//GetString 获取string
func (p *CmdData) GetString() (string, error) {
	num := len(p.data)
	if num < p.index+2 {
		return "", fmt.Errorf("%s", "string bytes lenth is invaild!")
	}
	len, err := BytesToInt(p.data[p.index:p.index+2], false)
	if err != nil {
		return "", fmt.Errorf("%s", "string bytes lenth is invaild!")
	}
	p.index += 2
	if num < p.index+len {
		return "", fmt.Errorf("%s", "string bytes lenth is invaild!")
	}
	s := string(p.data[p.index : p.index+len])
	p.index += len
	return s, nil
}

//GetBytes 获取Bytes
func (p *CmdData) GetBytes() ([]byte, error) {
	num := len(p.data)
	if num < p.index+2 {
		return nil, fmt.Errorf("%s", "string bytes lenth is invaild!")
	}
	len, err := BytesToInt(p.data[p.index:p.index+2], false)
	if err != nil {
		return nil, fmt.Errorf("%s", "string bytes lenth is invaild!")
	}
	p.index += 2
	if num < p.index+len {
		return nil, fmt.Errorf("%s", "string bytes lenth is invaild!")
	}
	s := p.data[p.index : p.index+len]
	p.index += len
	return s, nil
}

//AddCmdID 添加CmdId
func (p *CmdData) AddCmdID(num int) bool {
	return p.AddInt16(num)
}

//AddInt 添加int
func (p *CmdData) AddInt(num int) bool {
	data, err := IntToBytes(num, 4)
	if err != nil {
		return false
	}
	p.data = BytesCombine(p.data, data)
	return true
}

//AddInt8 添加int8
func (p *CmdData) AddInt8(num int) bool {
	data, err := IntToBytes(num, 1)
	if err != nil {
		return false
	}
	p.data = BytesCombine(p.data, data)
	return true
}

//AddInt16 添加int16
func (p *CmdData) AddInt16(num int) bool {
	data, err := IntToBytes(num, 2)
	if err != nil {
		return false
	}
	p.data = BytesCombine(p.data, data)
	return true
}

//AddInt64 添加int64
func (p *CmdData) AddInt64(num int64) bool {
	data, err := IntToBytes64(num)
	if err != nil {
		return false
	}
	p.data = BytesCombine(p.data, data)
	return true
}

//AddString 添加string
func (p *CmdData) AddString(s string) bool {
	len := len([]rune(s))
	lendata, err := IntToBytes(len, 2)
	if err != nil {
		return false
	}
	p.data = BytesCombine(p.data, lendata, []byte(s))
	return true
}

//AddBytes AddBytes
func (p *CmdData) AddBytes(b []byte) bool {
	len := len(b)
	lendata, err := IntToBytes(len, 2)
	if err != nil {
		return false
	}
	p.data = BytesCombine(p.data, lendata, b)
	return true
}

//Data 获取发送包
func (p *CmdData) Data() []byte {
	len := len(p.data) + 2
	lendata, err := IntToBytes(len, 2)
	if err != nil {
		return nil
	}
	return BytesCombine(lendata, p.data)
}
