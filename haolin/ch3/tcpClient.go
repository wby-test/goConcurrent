package main

import (
	"bytes"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	SERVER_NETWORK = "tcp"
	SERVER_ADDRESS = "127.0.0.1:8085"
	DELIMITER      = '\t'
)

var wg sync.WaitGroup

func printLog(role string, sn int, format string, args ...interface{}) {
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Printf("%s[%d]: %s", role, sn, fmt.Sprintf(format, args...))
}

func printServerlog(format string, args ...interface{}) {
	printLog("Server", 0, format, args)
}

func printClientlog(sn int, format string, args ...interface{}) {
	printLog("Client", sn, format, args)
}

func read(conn net.Conn) (string, error) {
	readBytes := make([]byte, 1)
	var buffer bytes.Buffer
	for {
		_, err := conn.Read(readBytes)
		if err != nil {
			return "", err
		}
		readByte := readBytes[0]
		if readByte == DELIMITER {
			break
		}
		buffer.WriteByte(readByte)
	}
	return buffer.String(), nil
}

func write(conn net.Conn, content string) (int, error) {
	var buffer bytes.Buffer
	buffer.WriteString(content)
	buffer.WriteByte(DELIMITER)
	return conn.Write(buffer.Bytes())
}

func handleCoon(conn net.Conn) {
	defer conn.Close()
	for {
		conn.SetDeadline(time.Now().Add(10 * time.Second))
		strReq, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printServerlog("the connection is closed by other side.")
			} else {
				printServerlog("read error: %s", err)
			}
			break
		}
		printServerlog("receive request: %s", strReq)

		intReq, err := strToInt32(strReq)
		if err != nil {
			n, err := write(conn, err.Error())
			printServerlog("sent error message(written %d bytes): %s", n, err)
		}

		floatResp := math.Cbrt(float64(intReq))
		respMsg := fmt.Sprintf("the cube root of %d is %f.", intReq, floatResp)
		n, err := write(conn, respMsg)
		if err != nil {
			printServerlog("Write error: %s", err)
		}
		printServerlog("sent response (written %d bytes) : %s.", n, respMsg)
	}
}

func strToInt32(str string) (int32, error) {
	val, err := strconv.Atoi(str)
	if err != nil {
		panic("strconv false")
	}
	return int32(val), err
}

func serverGo() {
	defer wg.Done()
	var listener net.Listener
	listener, err := net.Listen(SERVER_NETWORK, SERVER_ADDRESS)
	if err != nil {
		printServerlog("listen Error: %s", err)
	}

	printServerlog("got listener for the server:(local address: %s)", listener.Addr())
	for {
		conn, err := listener.Accept()
		if err != nil {
			printServerlog("accept error: %s", err)
		}
		printServerlog("established a connection with a client application. (remote address: %s)", conn.RemoteAddr())
		go handleCoon(conn)
	}
}

func clientGo(id int) {
	defer wg.Done()
	conn, err := net.DialTimeout(SERVER_NETWORK, SERVER_ADDRESS, 2*time.Second)
	if err != nil {
		printClientlog(id, "dial error : %s", err)
		return
	}

	defer conn.Close()
	printClientlog(id, "connected to server.(remote address: %s, local address : %s)", conn.RemoteAddr(), conn.LocalAddr())
	time.Sleep(200 * time.Millisecond)

	requestNum := 5
	conn.SetDeadline(time.Now().Add(5 * time.Millisecond))
	for i := 0; i < requestNum; i++ {
		req := rand.Int31()
		n, err := write(conn, fmt.Sprintf("%d", req))
		if err != nil {
			printClientlog(id, "write error", err)
			continue
		}
		printClientlog(id, "send request (written %d bytes): %d.", n, req)
	}

	for i := 0; i < requestNum; i++ {
		strRsp, err := read(conn)
		if err != nil {
			if err == io.EOF {
				printClientlog(id, "the connection is closed by another sid")
			} else {
				printClientlog(id, "read Error. %s", err)
			}
			break
		}
		printClientlog(id, "received response: %s.", strRsp)
	}

}

func main() {
	wg.Add(2)
	go serverGo()
	time.Sleep(1 * time.Second)
	go clientGo(1)
	wg.Wait()
}
