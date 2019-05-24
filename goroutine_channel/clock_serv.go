// 并发时钟服务器：它以每秒钟一次的频率向客户端发送当前时间

package main

import (
    "io"
    "log"
    "net"
    "time"
)

func main() {
    // 创建一个监听对象，它在端口上监听已到达的连接
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }

    for {
        // 返回一个net.Conn对象，表示一个已接受的新连接对象
        conn, err := listener.Accept()  // 阻塞：等待客户端连接的到达
        if err != nil {
            log.Print(err)  // 异常：例如连接终止
            continue
        }
        // 此实现只支持顺序执行，当有一个客户端连接时，后续的连接只有
        // 等此连接终止后才能被执行
        // handleConn(conn)  // 处理已接受的连接
        // 支持并发连接的版本
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()  // 延迟调用
    for {
        // 向客户端发送当前时间
        _, err := io.WriteString(conn, time.Now().Format("15:04:05\n"))
        if err != nil {
            return  // 错误：如客户端关闭连接等
        }
        time.Sleep(1 * time.Second)  // 等待一秒钟
    }
}
