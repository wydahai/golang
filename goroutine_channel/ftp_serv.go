// 实现一个简单的FTP服务器

package main

import (
    "fmt"
    "log"
    "net"
    "flag"
    "io"
)

var port = flag.String("port", "8000", "port number to listen on")

func main() {
    flag.Parse()
    fmt.Printf("FTP server will listen on %s\n", *port)

    listener, err := net.Listen("tcp", "localhost:"+*port)
    if err != nil {
        log.Fatal(err)
    }

    for {  // 等待客户端的连接到达
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("accept a new conn error: %s", err)
        }

        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    buffer := make([]byte, 1024)
    clientAddr := conn.RemoteAddr().String()
    for {  // 等待处理客户端的命令
        n, err := conn.Read(buffer)
        if err != nil {
            if err == io.EOF {
                log.Printf("connection: %s closed", clientAddr)
            } else {
                log.Printf("%s: connetion error: %s", clientAddr, err)
            }
            return
        }
        log.Printf("client: %s\tcommand: %s", clientAddr, string(buffer[:n]))
    }
}
