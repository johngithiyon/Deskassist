package main

import (
	"log"

	"github.com/johngithiyon/Deskassist/internal/services"
)

func main() {

	cmd,arg,_ := services.GetInput()

	log.Println(cmd)
	log.Println(arg)

}