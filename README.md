# netcmd
网络封装包

封装了接包和粘包,链接和接收消息回调
封装了CmdData数据包

NewListen 新建一个侦听接口
参数1:IP
参数2:端口
返回:无


CmdDial 链接服务器 
参数1 IP
参数2 端口
参数3 回调消息ID
返回:无

AddCmdData 注册消息
参数1 消息ID
参数2 回调函数
返回:无

CmdData接口:
AddCmdID 添加消息ID,接收时需要注册这个消息
参数1:16位消息ID

AddInt 添加一个int值
参数1:int
返回:bool 是否添加成功

AddInt8  添加一个int8值
参数1:int
返回:bool 是否添加成功

AddInt16 添加一个int16值
参数1:int
返回:bool 是否添加成功

AddInt64 添加一个int64值
参数1:int64
返回:bool 是否添加成功

AddString 添加一个string
参数1:添加一个string
返回:bool 是否添加成功

AddBytes 添加一个byte数组
参数1:[]byte
返回:bool 是否添加成功

Data() 获取所有添加数据的打包byte数组
返回:[]byte

GetInt 获取一个int
返回:int

GetInt8 获取一个int8
返回:int

GetInt16 获取一个int16
返回:int

GetInt64 获取一个int64
返回:int64

GetString 获取一个string
返回:string

GetBytes 获取[]byte
返回:[]byte
