package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)
	//Go 实现并发 TCP 端口扫描器
//这里通过互斥锁来解决数据竞争问题，使用WaitGroup来解决协程同步的问题，TCPScanner代码如下：
func main() {
	var wg sync.WaitGroup
	var mutex sync.Mutex
	ports := make([]int, 0)
	for i := 80; i <= 50000; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			conn, err := net.DialTimeout("tcp", fmt.Sprintf("127.0.0.1:%d", port), time.Second)
			if err != nil {
				log.Printf("Error:%v.Port:[%d]\n", err, port)
			} else {
				conn.Close()
				log.Printf("Connection successful.Port:[%d]\n", port)
				mutex.Lock()
				ports = append(ports, port)
				mutex.Unlock()
			}
		}(i)
	}
	wg.Wait()
	fmt.Printf("Opened ports:%v", ports)
}

