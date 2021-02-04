package main

import (
    "fmt"
    "log"
    "net"
    "time"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:9090")
    if err != nil {
        log.Fatalln(err)
    }

    _, err = conn.Write([]byte("Hi Server!"))
    time.Sleep(500 * time.Millisecond)
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Message sent: Hi Server!")

    _, err = conn.Write([]byte("How are you?"))
    if err != nil {
        log.Fatalln(err)
    }
    fmt.Println("Message sent: How are you?")

    for {
        buffer := make([]byte, 1400)
        dataSize, err := conn.Read(buffer)
        if err != nil {
            fmt.Println("connection closed")
            return
        }

        data := buffer[:dataSize]
        fmt.Println("received message: ", string(data))
    }
} 
