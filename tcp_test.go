// @Title  
// @Description  
// @Author  Wangwengang  2022/4/30 下午12:12
// @Update  Wangwengang  2022/4/30 下午12:12
package main

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/wwengg/arsenal/anet/impl"
)

func TestTcp(t *testing.T){
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}

	for {
		fmt.Println("发送数据")
		dp := impl.NewDataPack()
		msg := impl.NewMsgPackage(0,nil)
		bytes,_ := dp.Pack(msg)
		_, err := conn.Write(bytes)
		if err != nil {
			fmt.Println("write error err ", err)
			return
		}

		time.Sleep(5 * time.Second)
	}
}
