package main

import (
    "bufio"
    "fmt"
    "net"
    "strings"
    "sync"    //"time"
)

type customMap struct { 
    sync.RWMutex
    store map[string]string 
}

var m customMap

//var lock = sync.RWMutex{}

const PORT = "5555"

func main() {
    m = customMap { store: make(map[string]string) }
    fmt.Println("starting server on " + PORT)
    //var hashMap = make(map[string]interface{}) fmt.Println(hashMap)
    socket, err := net.Listen("tcp", ":"+PORT)
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
        processData(&data, conn)
    }
}

/* func printMap() { for { time.Sleep( 5 * time.Second ) fmt.Println(hashMap) } }*/

func processData(data *[]byte, conn net.Conn) {
    //fmt.Println(string(*data))
    command := string(*data)
    splitted := strings.Split(command, " ")
    switch splitted[0] {
    case "GET":
        if len(splitted) < 2 {
            conn.Write([]byte("Usage: GET key \n"))
            break
        }
        val := getCommand(splitted[1])
        // fmt.Println(val)
        conn.Write([]byte(val + "\n"))
        break
    case "SET":
        if len(splitted) < 3 {
            conn.Write([]byte("Usage: SET key value \n"))
            break
        }
        setCommand(splitted[1], splitted[2])
        conn.Write([]byte("done\n"))
        break
    case "DEL":
        if len(splitted) < 2 {
            conn.Write([]byte("Usage: DEL key \n"))
            break
        }
        delCommand(splitted[1])
        conn.Write([]byte("done\n"))
        break
    default:
        conn.Write([]byte("Only GET & SET implemented\n"))
        break
    }
}

func getCommand(key string) string {
    key = strings.TrimSpace(key)
    m.RLock()
    var v string = m.store[key]
    m.RUnlock()
    return v
}

func setCommand(key string, val string) {
    // fmt.Println(key, val)
    m.RLock()
    m.store[key] = val
    m.RUnlock()
}

func delCommand(key string) {
    key = strings.TrimSpace(key)
    _, ok := m.store[key]
    if ok {
        delete(m.store, key)
    }
}


