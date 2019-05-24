// 模拟时钟墙：同时显示多个地区的时间

package main

import (
    "io"
    "net"
    "os"
    "fmt"
    "strings"
    "time"
)

func main() {
    for _, args := range os.Args[1:] {
        addr := strings.Split(args, "=")[1]
        go getTime(addr)
    }

    for {
        time.Sleep(10 * time.Second)
    }
}

func getTime(addr string) {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println(err)
        return
    }
    defer conn.Close()

    mustCopy(os.Stdout, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
    if _, err := io.Copy(dst, src); err != nil {
        return
    }
}
