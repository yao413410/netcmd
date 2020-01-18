package netcmd

import (
	"net"
	"strconv"
	"sync"
	"time"
)

const (
	//NETERROR 掉线,退出消息
	NETERROR = -1
	//NETDIALOK 链接服务器成功消息,因为可能会需要连接几个服务器,所以多接几个ID
	NETDIALOK_1 = -2
	NETDIALOK_2 = -3
	NETDIALOK_3 = -4
	NETDIALOK_4 = -5
	NETDIALOK_5 = -6
)

//NetFuncData 接口回调函数
type NetFuncData func(net.Conn, *CmdData) error

//CmdListData NetFuncData数组
var CmdListData map[int]NetFuncData

//AddCmdData 注册消息
func AddCmdData(id int, function NetFuncData) {
	if CmdListData == nil {
		CmdListData = make(map[int]NetFuncData)
	}
	CmdListData[id] = function
}

//NowTime 获取当前时间
func NowTime() string {
	return time.Now().String()
}

//NewListen 新建一个侦听接口
func NewListen(ip string, port int) {
	go INewListen(ip, port)
}

//INewListen 并发侦听接口
func INewListen(ip string, port int) {
	listen, err := net.Listen("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		PrintfWarning(NowTime(), "listen error", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			PrintfWarning(NowTime(), "Accept err = ", err)
		}
		go IProcess(conn)
	}
}

var pCmdLock sync.Mutex //互斥锁
//IProcess 并发消息接口
func IProcess(conn net.Conn) {
	defer conn.Close()
	var iReadType int = 0
	var bytes []byte
	for {
		buf := make([]byte, 256)

		//fmt.Println(NowTime(), "server waiting for client: "+conn.RemoteAddr().String())
		n, err := conn.Read(buf)
		if err != nil {
			//fmt.Println(time.Now().String(), "server read err: ", err)
			if CmdListData != nil {
				pDataFuc, ok := CmdListData[NETERROR]
				if ok {
					err := pDataFuc(conn, nil)
					if err != nil {
						PrintfWarning("Cmd Analysis cmd id=%d ,%t", NETERROR, err)
					}
				}
			}
			return
		}

		b := buf[:n]
		//fmt.Println("bytes=", b)
		if iReadType == -1 {
			bytes = BytesCombine(bytes, b)
		} else {
			bytes = b
		}
		iReadType = CmdByte(conn, bytes)
		for {
			if iReadType > 0 {
				bytes = bytes[iReadType:]
				iReadType = CmdByte(conn, bytes)
			} else {
				break
			}
		}
		//fmt.Println(NowTime(), "server waiting for client: type", iReadType, bytes)
		buf = make([]byte, 256)
	}
}

//CmdByte 接包和粘包
func CmdByte(conn net.Conn, bytes []byte) int {
	nlen, err := BytesToInt(bytes[:2], true)
	//fmt.Println("nlen=", nlen)
	if err != nil {
		return 0
	}
	n := len(bytes)
	if nlen == n {
		CmdAnalysis(conn, bytes[2:])
		return 0
	} else if nlen < n {
		bytes = bytes[:nlen]
		CmdAnalysis(conn, bytes[2:])
		return nlen
	} else {
		return -1
	}
}

//CmdAnalysis 解析包
func CmdAnalysis(conn net.Conn, bytes []byte) {
	if len(bytes) < 2 {
		return
	}
	cmdID, err := BytesToInt(bytes[:2], true)
	if err != nil {
		return
	}
	pCmdLock.Lock()
	if CmdListData != nil {
		pDataFuc, ok := CmdListData[cmdID]
		if ok {
			var data CmdData
			data.InitData(bytes[2:])
			err := pDataFuc(conn, &data)
			if err != nil {
				PrintfWarning("Cmd Analysis cmd id=%d ,%t", cmdID, err)
			}
		} else {
			PrintfWarning("Cmd Analysis no cmd id=%d", cmdID)
		}
	}
	pCmdLock.Unlock()
}

//CmdDial 链接服务器 ,cmdid为回调函数的id
func CmdDial(ip string, port int, cmdid int) {
	go ICmdDial(ip, port, cmdid)
}

//ICmdDial 链接服务器
func ICmdDial(ip string, port int, cmdid int) {
	conn, err := net.Dial("tcp", ip+":"+strconv.Itoa(port))
	if err != nil {
		if CmdListData != nil {
			pDataFuc, ok := CmdListData[NETERROR]
			if ok {
				err := pDataFuc(conn, nil)
				if err != nil {
					PrintfWarning("Dial error %t", NETERROR, err)
				}
			}
		}
		return
	}
	defer conn.Close()
	if CmdListData != nil {
		pDataFuc, ok := CmdListData[cmdid]
		if ok {
			var data CmdData
			data.AddString("dial server " + ip + " ok")
			err := pDataFuc(conn, &data)
			if err != nil {
				PrintfWarning("Cmd Analysis cmd id=%d ,%t", cmdid, err)
			}
		} else {
			PrintfWarning("Cmd Analysis no cmd id=%d", cmdid)
		}
	}
	var iReadType int = 0
	var bytes []byte
	for {
		buf := make([]byte, 256)
		n, err := conn.Read(buf)
		if err != nil {
			if CmdListData != nil {
				pDataFuc, ok := CmdListData[NETERROR]
				if ok {
					err := pDataFuc(conn, nil)
					if err != nil {
						PrintfWarning("Cmd Analysis cmd id=%d ,%t", NETERROR, err)
					}
				}
			}
			return
		}

		b := buf[:n]
		if iReadType == -1 {
			bytes = BytesCombine(bytes, b)
		} else {
			bytes = b
		}
		iReadType = CmdByte(conn, bytes)
		for {
			if iReadType > 0 {
				bytes = bytes[iReadType:]
				iReadType = CmdByte(conn, bytes)
			} else {
				break
			}
		}
		buf = make([]byte, 256)
	}
}
