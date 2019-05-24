// 回显服务器

package main

import (
    "fmt"
    "log"
    "net"
    "time"
    "bufio"
    "strings"
)

func main() {
    listener, err := net.Listen("tcp", "localhost:8000")
    if err != nil {
        log.Fatal(err)
    }
    defer listener.Close()

    for {
        conn, err := listener.Accept()
        if err != nil {
            log.Printf("accept error: %s", err)
            continue
        }
        go handleConn(conn)
    }
}

func handleConn(conn net.Conn) {
    defer conn.Close()

    input := bufio.NewScanner(conn)
    for input.Scan() {
        go echo(conn, input.Text(), 3*time.Second)
    }
}

func echo(conn net.Conn, msg string, delay time.Duration) {
    fmt.Fprintln(conn, "\t", strings.ToUpper(msg))
    time.Sleep(delay)
    fmt.Fprintln(conn, "\t", msg)
    time.Sleep(delay)
    fmt.Fprintln(conn, "\t", strings.ToLower(msg))
}
