package main

import (
	"crypto/tls"
	"fmt"
	"time"
)

type Config struct {
	Protocol string
	Timeout  time.Duration
	Maxconns int
	TLS      *tls.Config
}

type Server struct {
	Addr string
	Port int
	Conf *Config
}

func NewServer(addr string, port int, conf *Config) (*Server, error) {
	if conf == nil {
		conf = &Config{}
	}
	return &Server{Addr: addr, Port: port, Conf: conf}, nil
}

func main() {
	//Using the default configuratrion
	srv1, _ := NewServer("localhost", 9000, nil)
	fmt.Println(srv1)

	conf := Config{Protocol: "tcp", Timeout: 60 * time.Second}
	srv2, _ := NewServer("locahost", 9000, &conf)
	fmt.Println(srv2)

	s := map[interface{}]interface{}{"test": "test", "a": 1, "b": 2, "c": true}
	fmt.Println(s)
}
