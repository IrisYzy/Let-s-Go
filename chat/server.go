package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main()  {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil{
		//log.Fatal()打印错误信息并调用os.Exit(1),终止程序
		log.Fatal(err)
	}

	//
	go broadcaster()

	for  {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		//每个客户端一个goroutine
		go handleConn(conn)
	}
}

type client chan <- string

var (
	entering = make(chan client)
	leaving = make(chan client)
	message = make(chan string)
)

func broadcaster()  {
	clients := make(map[client]bool)
	for  {
		select {
		case msg := <- message:
			for cli := range clients{
				cli <- msg
			}
		case cli := <- entering:
			clients[cli] = true
		case cli := <-leaving:
			delete(clients, cli)
			close(cli)
		}
	}
}

func handleConn(conn net.Conn)  {
	ch := make(chan string)

	//写入消息到客户端的连接
	go writeToClient(conn,ch)

	who := conn.RemoteAddr().String()

	ch <- "You are " + who

	message <- who + "are arrived"

	entering <- ch

	input := bufio.NewScanner(conn)

	//阻塞监听客户端输入
	for input.Scan() {
		message <- who + ": " + input.Text()
	}

	//客户端断开连接
	leaving <- ch
	message <- who + "are left"
	conn.Close()
}

func writeToClient(conn net.Conn, ch <- chan string)  {
	for msg := range ch{
		fmt.Println(conn, msg)
	}
}


