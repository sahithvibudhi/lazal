package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    //"time"
)

var hashMap = make(map[string]string)

func main() {
    fmt.Println("starting server on 8080")
    //var hashMap = make(map[string]interface{}) fmt.Println(hashMap)
    socket, err := net.Listen("tcp", ":8080")
    if err != nil {
        panic(err)
    }
    for {
        conn, err := socket.Accept()
        if err != nil {
            panic(err)
        }
        go handleConnection(conn)
        //go printMap()
    }
}

func handleConnection(conn net.Conn) {
    for {
        data, err := bufio.NewReader(conn).ReadBytes('\r')
        if err != nil {
            // Handle error
        }
        go processData(&data, conn)
    }
}

/* func printMap() { for { time.Sleep( 5 * time.Second ) fmt.Println(hashMap) } }*/

func processData(data *[]byte, conn net.Conn) {
    //fmt.Println(string(*data))
    command := string(*data)
    splitted := strings.Split(command, " ")
    switch splitted[0] {
    case "GET":
        val := getCommand(splitted[1])
        // fmt.Println(val)
        conn.Write([]byte(val + "\n"))
        break
    case "SET":
        setCommand(splitted[1], splitted[2])
        conn.Write([]byte("done\n"))
        break
    default:
        conn.Write([]byte("Only GET & SET implemented\n"))
        break
    }
}

func getCommand(key string) string {
    key = strings.TrimSpace(key)
    return hashMap[key]
}

func setCommand(key string, val string) {
    // fmt.Println(key, val)
    hashMap[key] = val
}




