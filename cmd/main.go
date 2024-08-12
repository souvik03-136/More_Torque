package main

import "github.com/souvik03-136/more-torque/internal/server"

func main() {
	srv := server.NewServer()
	if err := srv.Run(); err != nil {
		panic(err)
	}
}
