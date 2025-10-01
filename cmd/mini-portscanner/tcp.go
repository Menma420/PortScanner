package main

import (
	"fmt"
	"net"
	"time"
)

func checkTCP(ip string, port int, timeout time.Duration) (bool, time.Duration, error){
	// 	Build address string: addr
	addr := fmt.Sprintf("%s:%d", ip, port)

	// Record start time: start := time.Now().
	start := time.Now()

	// Use net.DialTimeout("tcp", addr, timeout) to attempt the connection.
	conn, err := net.DialTimeout("tcp", addr, timeout)

	// Compute latency := time.Since(start).
	latency := time.Since(start)

	// If DialTimeout returned a conn and err == nil:
	// Immediately conn.Close() (we only needed to know it opened).
	if err == nil{
		defer conn.Close()
		fmt.Println("open: true  latency:"+latency.String()+"  err: <nil>") 
		return true, latency, nil
	}else{
		fmt.Println("open: false  latency:"+latency.String()+"  err: "+err.Error())
		return false, latency, err
	}

}