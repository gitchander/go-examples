package main

import (
	"fmt"
	"time"
)

type Server struct{ quit chan bool }

func OpenServer() *Server {

	s := &Server{make(chan bool)}
	go s.run()
	return s
}

func (s *Server) run() {

	ticker := time.NewTicker(time.Millisecond * 1000)
	for {
		select {
		case <-s.quit:
			{
				fmt.Println("final task")
				time.Sleep(time.Second)
				fmt.Println("task done")

				s.quit <- true
				return
			}

		case <-ticker.C:
			fmt.Println("time to work")
		}
	}
}

func (s *Server) Close() {

	if s.quit != nil {

		s.quit <- true
		<-s.quit

		close(s.quit)
		s.quit = nil
	}
}
