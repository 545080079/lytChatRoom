package utils

import (
	"awesomeProject/YTChat/Common"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)


type Transfer struct{
	Conn net.Conn
	Buf [8096]byte	//切片 = 数组，传输时使用的缓冲
}

func (t *Transfer) ReadPkg() (message Common.Message, err error){

	//fmt.Println("客户端0： 监听服务端发送的消息....")
	cnt, err := t.Conn.Read(t.Buf[:4])
	if cnt != 4 || err != nil{
		fmt.Println("conn.Read Size Failed: ", err)
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(t.Buf[0:4])

	cnt, err = t.Conn.Read(t.Buf[:pkgLen])
	if cnt != int(pkgLen) || err != nil{
		fmt.Println("conn.Read failed:", err)
		return
	}

	//这里message一定要取它地址
	err = json.Unmarshal(t.Buf[:pkgLen], &message)
	if err != nil{
		fmt.Println("json.Unmarshal failed: ", err)
		return
	}
	return message, err
}



func (t *Transfer)WritePkg(data []byte) (err error){
	//先发送长度
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(t.Buf[0:4], pkgLen)
	cnt, err := t.Conn.Write(t.Buf[:4])
	if cnt != 4 || err != nil{
		fmt.Println("conn.write_cnt failed: ", err)
		return
	}


	//发送data本身
	_, err = t.Conn.Write(data)
	if err != nil{
		fmt.Println("conn.write_data failed: ", err)
		return err
	}

	return
}