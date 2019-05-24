package main

import (
    "io"
    "net"
    "log"
    "time"
    "flag"
    "fmt"
)

var port = flag.String("port", "8000", "port number")
var region = flag.String("region", "unknown", "region name")

func main() {
    flag.Parse()
    fmt.Printf("listen on: %s\n", *port)

    listener, err := net.Listen("tcp", "localhost:"+*port)
    if err != nil {
        log.Fatal(err)
    }

    for {  // 循环监听连接到达并进行处理
        conn, err := listener.Accept()
        if err != nil {
            log.Print(err)
            continue
        }

        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()
    for {  // 循环向客户端发送当前时间
        msg := *region + ": " + time.Now().Format("16:19:20\n")
        _, err := io.WriteString(conn, msg)
        if err != nil {
            return
        }
        time.Sleep(1 * time.Second)
    }
}
