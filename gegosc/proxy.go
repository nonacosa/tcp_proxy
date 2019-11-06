package gegosc

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"net"
	"os"
	"strconv"
	"time"
)

var ListenPorts = make(map[int]string)

func ProxyStart(fromPort, toPort int, ip string) {

	proxyAddr := fmt.Sprintf(":%d", fromPort)

	if _, ok := ListenPorts[fromPort]; !ok {

		proxyListener, err := net.Listen("tcp", proxyAddr)
		if err != nil {
			fmt.Println("Unable to listen on: %s, error: %s\n", proxyAddr, err.Error())
			os.Exit(1)
		} else {
			ListenPorts[fromPort] = proxyAddr
		}
		defer proxyListener.Close()

		for {
			proxyConn, err := proxyListener.Accept()
			if err != nil {
				fmt.Printf("Unable to accept a request, error: %s\n", err.Error())
				continue
			}

			// Read a header firstly in case you could have opportunity to check request
			// whether to decline or proceed the request
			buffer := make([]byte, 1024)
			n, err := proxyConn.Read(buffer)
			if err != nil {
				fmt.Printf("Unable to read from input, error: %s\n", err.Error())
				continue
			}

			// TODO
			// Your choice to make decision based on request header 172.24.167.144:9200

			targetAddr := fmt.Sprintf("%s:%d", ip, toPort);
			targetConn, err := net.Dial("tcp", targetAddr)
			if err != nil {
				fmt.Println("Unable to connect to: %s, error: %s\n", targetAddr, err.Error())
				proxyConn.Close()
				continue
			}

			n, err = targetConn.Write(buffer[:n])
			if err != nil {
				fmt.Printf("Unable to write to output, error: %s\n", err.Error())
				proxyConn.Close()
				targetConn.Close()
				continue
			}

			//新建计时器，10秒后触发

			go func() {
				timer := time.NewTimer(time.Second * 20)

				for {
					// 多路复用通道
					select {
					case <-timer.C: // 计时器到时了
						fmt.Println("stop")
						closeProxy(proxyConn, targetConn, fromPort)
						// 跳出循环

					}

				}
			}()

			eachProxyRequest(proxyConn, targetConn)
		}

	} else {
		ProxyStart(fromPort+1, toPort, ip)
		//return errors.New("CheckPortType ERROR: this port is exist")
	}

}

func eachProxyRequest(r, w net.Conn) {
	// todo user:port 元祖 已经挂掉的 不准再打洞
	if len(ListenPorts) > 2 {
		r.Close()
		w.Close()
		print("权限不够关闭连接")
		delete(ListenPorts, 10221)
		return
	}
	//验证
	go proxyRequest(r, w)
	go proxyRequest(w, r)
}

func closeProxy(r, w net.Conn, port int) {
	r.Close()
	w.Close()

	if _, ok := ListenPorts[port]; ok {
		delete(ListenPorts, port)
		print("权限时间到，系统关闭 %d 连接", port)
	}

}

// Forward all requests from r to w
func proxyRequest(r, w net.Conn) {
	defer r.Close()
	defer w.Close()

	var buffer = make([]byte, 4096000)
	for {
		n, err := r.Read(buffer)
		if err != nil {
			fmt.Printf("Unable to read from input, error: %s\n", err.Error())
			break
		}

		n, err = w.Write(buffer[:n])
		if err != nil {
			fmt.Printf("Unable to write to output, error: %s\n", err.Error())
			break
		}
	}
}

func RandPort(min, max int64) int {
	maxBigInt := big.NewInt(max)
	i, _ := rand.Int(rand.Reader, maxBigInt)
	if i.Int64() < min {
		RandPort(min, max)
	}
	strInt64 := strconv.FormatInt(i.Int64(), 10)
	id16, _ := strconv.Atoi(strInt64)
	return id16
}
