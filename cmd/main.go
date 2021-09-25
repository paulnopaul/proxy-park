package main

import (
	"bufio"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net"
	"net/http"
	"syscall"
)

func PrepareRequest(r *http.Request) {
	r.RequestURI = ""
	r.Header.Del("Proxy-Connection")
}

func RequestProxy(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	log.Println("Proxy", r.RequestURI)
	client := http.Client{}

	for key, valueArr := range r.Header {
		for _, value := range valueArr {
			fmt.Println(key, value)
		}
	}

	resp, err := client.Do(r)
	if err != nil {
		log.Println("ERR")
	}

	for key, valueArr := range resp.Header {
		for _, value := range valueArr {
			w.Header().Add(key, value)
		}
	}
	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	_, _ = fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type Socket struct {
	fd int
}

func newTCPSocket(addr string, port int) (*Socket, error) {
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_STREAM, 0)
	if err != nil {
		return nil, err
	}
	sa := &syscall.SockaddrInet4{Port: port}
	copy(sa.Addr[:], net.ParseIP(addr))

	err = syscall.Bind(fd, sa)
	if err != nil {
		return nil, err
	}

	err = syscall.Listen(fd, syscall.SOMAXCONN)
	if err != nil {
		return nil, err
	}
	return &Socket{fd}, nil
}

func (s *Socket) Accept() (*Socket, error) {
	nfd, _, err := syscall.Accept(s.fd)
	if err != nil {
		return nil, err
	}
	syscall.CloseOnExec(nfd)
	return &Socket{nfd}, nil
}

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()

	fmt.Printf("%+v connected\n", name)
	conn.Write([]byte("Hello, " + name + "\n\r"))

	defer conn.Close()

	reader := bufio.NewReader(conn)
	rq, err := http.ReadRequest(reader)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rq.Method)
	fmt.Println(rq.Host)
}

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}
