
package main

import (
"net"
"log"
"io"
"os"
)

func main()  {
	conn, err := net.Dial("tcp","localhost:8000")
	if err != nil{
		log.Fatal(err)
	}
	defer conn.Close()

	//获取服务端消息
	go io.Copy(os.Stdout,conn)

	//将用户输入的文本消息发送到服务端
	io.Copy(conn, os.Stdin)
}

func ioCopy(dst io.Writer, src io.Reader)  {
	if _, err := io.Copy(dst,src); err != nil{
		log.Fatal(err)
	}
}
