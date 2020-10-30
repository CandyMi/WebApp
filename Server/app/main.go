package main

import (
	"app/service"
)

func main() {
	s := service.New()
	s.Run()
}
